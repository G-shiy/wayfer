package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/G-shiy/wayfer/internal/api/spec"
	"github.com/G-shiy/wayfer/internal/pgstore"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type store interface {
	GetParticipant(ctx context.Context, participantId uuid.UUID) (pgstore.Participant, error)
	ConfirmParticipant(ctx context.Context, participantId uuid.UUID) error
}

type ServerAPI struct {
	store  store
	logger *zap.Logger
}

func NewAPI(poll *pgxpool.Pool, logger *zap.Logger) ServerAPI {
	return ServerAPI{pgstore.New(poll), logger}
}

// Confirms a participant on a trip.
// (PATCH /participants/{participantId}/confirm)
func (serverapi *ServerAPI) PatchParticipantsParticipantIDConfirm(
	w http.ResponseWriter,
	r *http.Request,
	participantID string,
) *spec.Response {
	id, err := uuid.Parse(participantID)
	if err != nil {
		return spec.PatchParticipantsParticipantIDConfirmJSON400Response(spec.Error{Message: "uuid Inválido"})
	}

	participant, err := serverapi.store.GetParticipant(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return spec.PatchParticipantsParticipantIDConfirmJSON400Response(
				spec.Error{Message: "Participante não encontrado"},
			)
		}
		serverapi.logger.Error("Failed to get participant", zap.Error(err), zap.String("participant_id", participantID))
	}
	if participant.IsConfirmed {
		return spec.PatchParticipantsParticipantIDConfirmJSON400Response(
			spec.Error{Message: "participant já confirmado"},
		)
	}

	if err := serverapi.store.ConfirmParticipant(r.Context(), id); err != nil {
		serverapi.logger.Error(
			"Failt to confirm participant",
			zap.Error(err),
			zap.String("participant_id", participantID),
		)
		return spec.PatchParticipantsParticipantIDConfirmJSON400Response(
			spec.Error{Message: "something went wrong, try again"},
		)
	}
	return spec.PatchParticipantsParticipantIDConfirmJSON204Response(nil)
}

// Create a new trip
// (POST /trips)
func (serverapi *ServerAPI) PostTrips(w http.ResponseWriter, r *http.Request) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Get a trip details.
// (GET /trips/{tripId})
func (serverapi *ServerAPI) GetTripsTripID(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Update a trip.
// (PUT /trips/{tripId})
func (serverapi *ServerAPI) PutTripsTripID(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Get a trip activities.
// (GET /trips/{tripId}/activities)
func (serverapi *ServerAPI) GetTripsTripIDActivities(
	w http.ResponseWriter,
	r *http.Request,
	tripID string,
) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Create a trip activity.
// (POST /trips/{tripId}/activities)
func (serverapi *ServerAPI) PostTripsTripIDActivities(
	w http.ResponseWriter,
	r *http.Request,
	tripID string,
) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Confirm a trip and send e-mail invitations.
// (GET /trips/{tripId}/confirm)
func (serverapi *ServerAPI) GetTripsTripIDConfirm(
	w http.ResponseWriter,
	r *http.Request,
	tripID string,
) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Invite someone to the trip.
// (POST /trips/{tripId}/invites)
func (serverapi *ServerAPI) PostTripsTripIDInvites(
	w http.ResponseWriter,
	r *http.Request,
	tripID string,
) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Get a trip links.
// (GET /trips/{tripId}/links)
func (serverapi *ServerAPI) GetTripsTripIDLinks(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Create a trip link.
// (POST /trips/{tripId}/links)
func (serverapi *ServerAPI) PostTripsTripIDLinks(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Get a trip participants.
// (GET /trips/{tripId}/participants)
func (serverapi *ServerAPI) GetTripsTripIDParticipants(
	w http.ResponseWriter,
	r *http.Request,
	tripID string,
) *spec.Response {
	panic("not implemented") // TODO: Implement
}
