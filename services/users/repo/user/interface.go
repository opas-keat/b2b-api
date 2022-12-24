package user

import (
	"context"
	"users/entities"
)

//go:generate mockery --name Repo
type Repo interface {
	Create(ctx context.Context, member entities.User) (*entities.User, error)
	// List(ctx context.Context, filter entities.MemberFilter, pagination *models.Pagination, count *models.Count) (*[]entities.Member, int64, error)

	// Get(ctx context.Context, filter entities.Member) (*entities.Member, error)
	// GetMembersByMemberIDs(ctx context.Context, memberIDs []string, filter entities.Member) (*[]entities.Member, error)
	// ListReferrerMember(ctx context.Context, filter entities.MemberFilter, pagination *models.Pagination, count *models.Count, refIDs []string) (*[]entities.Member, int64, error)
	// Update(ctx context.Context, memberId string, updateModel entities.Member) error
	// GetRandomAgent(ctx context.Context, filter entities.GetRandomAgentFilter) (*entities.Member, error)
	// CreateMonitoring(ctx context.Context, monitoring entities.MMemberMonitoring) (*entities.MMemberMonitoring, error)
	// GetMonitoring(ctx context.Context, filter entities.MMemberMonitoring) (*entities.MMemberMonitoring, error)
	// UpdateMonitoring(ctx context.Context, id uint, updateModel entities.MMemberMonitoring) error
	// DeleteMonitoring(ctx context.Context, id uint) error
	// ListMonitoring(ctx context.Context, pagination *models.Pagination, count *models.Count, filter entities.MonitoringFilter) (*[]entities.Monitoring, int64, error)
}
