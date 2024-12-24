package store

import (
	"context"
)

type IStatisticsStore[M any] interface {
	IBaseStore[M]
	SumByGroup(ctx context.Context, group string, selectSQL string, conditions string, args ...interface{}) ([]*M, error)
	SumByGroupPage(ctx context.Context, group string, order interface{}, offset, limit int, selectSQL string, conditions string, args ...interface{}) ([]*M, int64, error)
}

type StatisticsStore[M any] struct {
	Store[M]
	db   IDB `autowired:""`
	name string
}

func (s *StatisticsStore[M]) OnComplete() {
	s.Store.OnComplete()
}

func (s *StatisticsStore[M]) SumByGroup(ctx context.Context, group string, selectSQL string, conditions string, args ...interface{}) ([]*M, error) {
	db := s.db.DB(ctx)
	results := make([]*M, 0)
	err := db.Model(s.Model).Select(selectSQL).Where(conditions, args...).Group(group).Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (s *StatisticsStore[M]) SumByGroupPage(ctx context.Context, group string, order interface{}, offset, limit int, selectSQL string, conditions string, args ...interface{}) ([]*M, int64, error) {
	db := s.db.DB(ctx)
	results := make([]*M, 0)
	var count int64
	err := db.Model(s.Model).Select(selectSQL).Where(conditions, args...).Group(group).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Model(s.Model).Select(selectSQL).Where(conditions, args...).Group(group).Order(order).Limit(limit).Offset(offset).Scan(&results).Error
	if err != nil {
		return nil, 0, err
	}
	return results, count, nil
}
