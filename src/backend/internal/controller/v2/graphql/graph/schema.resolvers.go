package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"time"

	"course/internal/controller/v2/graphql/graph/model"
	models "course/internal/model"
	"course/internal/service/dto"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, req model.RegisterRequest) (*model.RegisterResponse, error) {
	panic(fmt.Errorf("not implemented: Register - register"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// RefreshTokens is the resolver for the refreshTokens field.
func (r *mutationResolver) RefreshTokens(ctx context.Context, req model.RefreshTokensRequest) (*model.RefreshTokensResponse, error) {
	panic(fmt.Errorf("not implemented: RefreshTokens - refreshTokens"))
}

// FillProfile is the resolver for the fillProfile field.
func (r *mutationResolver) FillProfile(ctx context.Context, req model.FillProfileRequest) (*string, error) {
	panic(fmt.Errorf("not implemented: FillProfile - fillProfile"))
}

// ConfirmEmployeeInfoCard is the resolver for the confirmEmployeeInfoCard field.
func (r *mutationResolver) ConfirmEmployeeInfoCard(ctx context.Context, req model.ConfirmEmployeeInfoCardRequest) (*string, error) {
	panic(fmt.Errorf("not implemented: ConfirmEmployeeInfoCard - confirmEmployeeInfoCard"))
}

// CreatePassage is the resolver for the createPassage field.
func (r *mutationResolver) CreatePassage(ctx context.Context, req model.CreatePassageRequest) (*model.Passage, error) {
	passage, err := r.CheckpointService.CreatePassage(ctx, &dto.CreatePassageRequest{
		CheckpointID: 1,
		DocumentID:   models.ToDocumentID(int64(req.DocumentID)).Int(),
		Type:         models.ToPassageTypeFromString(req.Type).Int(),
		Time:         &req.Time,
		IsSQUID:      false,
	})
	if err != nil {
		return nil, err
	}

	return &model.Passage{
		ID:         int(passage.ID.Int()),
		DocumentID: int(passage.DocumentID.Int()),
		Type:       passage.Type.String(),
		Time:       *passage.Time,
		IsSquid:    passage.IsSQUID,
	}, err
}

// DeletePassage is the resolver for the deletePassage field.
func (r *mutationResolver) DeletePassage(ctx context.Context, req model.DeletePassageRequest) (string, error) {
	err := r.CheckpointService.DeletePassage(ctx, &dto.DeletePassageRequest{
		PassageID: int64(req.ID),
	})
	if err != nil {
		return "ERROR", err
	}

	return "OK", err
}

// CreateSQUIDPassage is the resolver for the createSQUIDPassage field.
func (r *mutationResolver) CreateSQUIDPassage(ctx context.Context, req model.CreateSQUIDPassageRequest) (*model.Passage, error) {
	passage, err := r.CheckpointService.CreatePassage(ctx, &dto.CreatePassageRequest{
		CheckpointID: 1,
		DocumentID:   models.ToDocumentID(int64(req.DocumentID)).Int(),
		Type:         models.ToPassageTypeFromString(req.Type).Int(),
		Time:         &req.Time,
		IsSQUID:      true,
	})
	if err != nil {
		return nil, err
	}

	return &model.Passage{
		ID:         int(passage.ID.Int()),
		DocumentID: int(passage.DocumentID.Int()),
		Type:       passage.Type.String(),
		Time:       *passage.Time,
		IsSquid:    passage.IsSQUID,
	}, err
}

// Healthcheck is the resolver for the healthcheck field.
func (r *queryResolver) Healthcheck(ctx context.Context) (*string, error) {
	now := time.Now().String()
	return &now, nil
}

// GetProfile is the resolver for the getProfile field.
func (r *queryResolver) GetProfile(ctx context.Context, req model.GetProfileRequest) (*model.GetProfileResponse, error) {
	panic(fmt.Errorf("not implemented: GetProfile - getProfile"))
}

// GetEmployeePhoto is the resolver for the getEmployeePhoto field.
func (r *queryResolver) GetEmployeePhoto(ctx context.Context, req model.GetEmployeePhotoRequest) (*string, error) {
	panic(fmt.Errorf("not implemented: GetEmployeePhoto - getEmployeePhoto"))
}

// ListFullInfoCards is the resolver for the listFullInfoCards field.
func (r *queryResolver) ListFullInfoCards(ctx context.Context, req model.ListFullInfoCardsRequest) (*model.ListFullInfoCardsResponse, error) {
	panic(fmt.Errorf("not implemented: ListFullInfoCards - listFullInfoCards"))
}

// GetFullInfoCard is the resolver for the getFullInfoCard field.
func (r *queryResolver) GetFullInfoCard(ctx context.Context, req model.GetFullInfoCardRequest) (*model.GetProfileResponse, error) {
	panic(fmt.Errorf("not implemented: GetFullInfoCard - getFullInfoCard"))
}

// GetEmployeeInfoCardPhoto is the resolver for the getEmployeeInfoCardPhoto field.
func (r *queryResolver) GetEmployeeInfoCardPhoto(ctx context.Context, req model.GetEmployeeInfoCardPhotoRequest) (*string, error) {
	panic(fmt.Errorf("not implemented: GetEmployeeInfoCardPhoto - getEmployeeInfoCardPhoto"))
}

// GetPassages is the resolver for the getPassages field.
func (r *queryResolver) GetPassages(ctx context.Context, req model.GetPassagesRequest) ([]*model.Passage, error) {
	passages, err := r.CheckpointService.ListPassages(ctx, &dto.ListPassagesRequest{
		DocumentID: int64(req.DocumentID),
	})
	if err != nil {
		return nil, err
	}

	modelPassages := make([]*model.Passage, len(passages))
	for i, passage := range passages {
		modelPassages[i] = &model.Passage{
			ID:         int(passage.ID.Int()),
			DocumentID: int(passage.DocumentID.Int()),
			Type:       passage.Type.String(),
			Time:       *passage.Time,
			IsSquid:    passage.IsSQUID,
		}
	}

	return modelPassages, err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
