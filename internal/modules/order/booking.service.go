package order

import (
	"go-api/internal/utils"

	profile "go-api/internal/modules/account"
	user "go-api/internal/modules/account"
	room "go-api/internal/modules/room"

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
	b.RoomID = room.RoomID

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

	return s.creadsteBookingInternal(b)
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
