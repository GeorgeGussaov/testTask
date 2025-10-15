package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"task/models"
)

var ErrNotFound = errors.New("subscription not found")

type SubscriptionsRepository struct {
	db *sql.DB
}

func NewSubscriptionsRepository(db *sql.DB) *SubscriptionsRepository {
	return &SubscriptionsRepository{db: db}
}

func (r *SubscriptionsRepository) SelectSumByInfo(ctx context.Context, queryInfo *models.SubscriptionFilter) (int, error) {
	sum := 0
	query := `SELECT SUM(price) AS total FROM subscriptions 
	WHERE start_date >= $1 AND start_date <= $2
	AND ($3::uuid IS NULL OR user_id = $3)
	AND ($4::text IS NULL OR service_name = $4);`
	err := r.db.QueryRowContext(ctx, query, queryInfo.FromDate, queryInfo.ToDate, queryInfo.UserID, queryInfo.ServiceName).Scan(&sum)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, ErrNotFound
		}
		return 0, err
	}
	return sum, nil
}

func (r *SubscriptionsRepository) SelectList(ctx context.Context) (subs []models.Subscription, err error) {
	query := `SELECT id, user_id, service_name, price, start_date
	FROM subscriptions
	ORDER BY start_date`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var subscription models.Subscription
		err := rows.Scan(
			&subscription.ID, &subscription.UserID, &subscription.ServiceName, &subscription.Price, &subscription.StartDate,
		)
		if err != nil {
			return nil, err
		}

		subs = append(subs, subscription)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return subs, nil
}

func (r *SubscriptionsRepository) Insert(ctx context.Context, subInfo models.Subscription) (string, error) {
	var id string
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO subscriptions (service_name, price, user_id, start_date)
		VALUES ($1, $2, $3, $4)
		RETURNING id`,
		subInfo.ServiceName,
		subInfo.Price,
		subInfo.UserID,
		subInfo.StartDate,
	).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("failed to insert subscription: %w", err)
	}

	return id, nil
}

func (r *SubscriptionsRepository) SelectById(ctx context.Context, id string) (*models.Subscription, error) {
	sub := &models.Subscription{}
	row := r.db.QueryRowContext(ctx, `SELECT id, user_id, service_name, price, start_date FROM subscriptions WHERE id = $1`, id)

	err := row.Scan(&sub.ID, &sub.UserID, &sub.ServiceName, &sub.Price, &sub.StartDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("err querying data from DB: %v", err)
	}
	return sub, nil
}

func (r *SubscriptionsRepository) Update(ctx context.Context, id string, sub *models.Subscription) error {
	res, err := r.db.ExecContext(ctx,
		`UPDATE subscriptions
	    SET service_name = $1, price = $2, user_id = $3, start_date = $4
	    WHERE id = $5`,
		sub.ServiceName,
		sub.Price,
		sub.UserID,
		sub.StartDate,
		id,
	)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *SubscriptionsRepository) DeleteById(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM subscriptions WHERE id = $1`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound
		}
		return fmt.Errorf("err querying data from DB: %v", err)
	}
	return nil
}
