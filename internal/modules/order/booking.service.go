package order

import (
	"go-api/internal/utils"

	profile "go-api/internal/modules/account"
	user "go-api/internal/modules/account"
	room "go-api/internal/modules/room"
	"time"

	"github.com/google/uuid"
)

type BookingService struct {
	BookingRepo *BookingRepository
	UserRepo    *user.UserRepository
	RoomRepo    *room.Repository
	ProfileRepo *profile.ProfileRepository
}

func NewRoleService(s *BookingRepository) *BookingService {
	return &BookingService{BookingRepo: s}
}

func (s *BookingService) CreateBooking(
	userUUID string,
	roomUUID string,
	b *Booking,
	guestName string,
	guestKTP string,
) error {

	user, err := s.UserRepo.FindByUuid(userUUID)
	if err != nil {
		return err
	}

	room, err := s.RoomRepo.FindByID(roomUUID)
	if err != nil {
		return err
	}

	b.UserID = user.ID
	b.RoomUuid = room.RoomID

	// 🔥 ambil profile
	profile, _ := s.ProfileRepo.FindByUserID(user.ID)

	if profile != nil && profile.FullName != "" && profile.KTP != "" {
		b.GuestName = profile.FullName
		b.GuestKTP = profile.KTP
	} else {
		if guestName == "" || guestKTP == "" {
			return utils.NewApiError(400, "name and KTP required")
		}

		b.GuestName = guestName
		b.GuestKTP = guestKTP
	}

	return s.createBookingInternal(b)
}

func (s *BookingService) generateAccess(b *Booking) error {
	tx, err := s.BookingRepo.DB.Begin()

	if err != nil {
		return err
	}
	defer tx.Rollback()

	token := uuid.New().String()

	return s.BookingRepo.CreateAccess(tx, &BookingAccess{
		BookingID:  b.ID,
		RoomID:     b.RoomID,
		QRCode:     token,
		ValidFrom:  *b.StartTime,
		ValidUntil: *b.EndTime,
		IsUsed:     true,
		AccessUUID: uuid.NewString(),
	})
}

func (s *BookingService) createBookingInternal(b *Booking) error {
	tx, err := s.BookingRepo.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 🔥 ambil room (pakai ID karena sudah di-set sebelumnya)
	room, err := s.RoomRepo.FindByID(b.RoomUuid)
	if err != nil {
		return err
	}

	// 🔥 hitung waktu + harga
	now := time.Now()

	switch b.BookingType {
	case "transit":
		if b.DurationHours == nil {
			return utils.NewApiError(400, "duration required")
		}

		end := now.Add(time.Duration(*b.DurationHours) * time.Hour)

		b.StartTime = &now
		b.EndTime = &end
		b.BasePrice = float64(*b.DurationHours) * room.PriceTransit

	case "daily":
		start := now
		end := now.Add(24 * time.Hour)

		b.StartTime = &start
		b.EndTime = &end
		b.BasePrice = room.PriceDaily

	case "monthly":
		start := now
		end := now.AddDate(0, 1, 0)

		b.StartTime = &start
		b.EndTime = &end
		b.BasePrice = room.PriceMonthly

	default:
		return utils.NewApiError(400, "invalid booking type")
	}

	if room.DepositAmount != nil {
		b.DepositAmount = *room.DepositAmount
	}

	b.Status = "pending"

	// 🔥 cek overlap
	exist, err := s.BookingRepo.CheckOverlap(tx, b.ID, *b.StartTime, *b.EndTime)
	if err != nil {
		return err
	}
	if exist {
		return utils.NewApiError(400, "room already booked")
	}

	// 🔥 create booking
	if err := s.BookingRepo.Create(tx, b); err != nil {
		return err
	}

	// 🔐 generate access (QR)
	if err := s.generateAccess(b); err != nil {
		return err
	}

	return tx.Commit()
}
