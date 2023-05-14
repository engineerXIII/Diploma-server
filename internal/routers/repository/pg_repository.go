package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/engineerXIII/Diploma-server/internal/models"
	"github.com/engineerXIII/Diploma-server/pkg/utils"
)

// Comments Repository
type routerRepo struct {
	db *sqlx.DB
}

// Comments Repository constructor
func NewrouterRepository(db *sqlx.DB) routers.Repository {
	return &routerRepo{db: db}
}

// Create comment
func (r *routerRepo) Create(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "routerRepo.Create")
	defer span.Finish()

	c := &models.Comment{}
	if err := r.db.QueryRowxContext(
		ctx,
		createComment,
		&comment.AuthorID,
		&comment.NewsID,
		&comment.Message,
	).StructScan(c); err != nil {
		return nil, errors.Wrap(err, "routerRepo.Create.StructScan")
	}

	return c, nil
}

// Update comment
func (r *routerRepo) Update(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "routerRepo.Update")
	defer span.Finish()

	comm := &models.Comment{}
	if err := r.db.QueryRowxContext(ctx, updateComment, comment.Message, comment.CommentID).StructScan(comm); err != nil {
		return nil, errors.Wrap(err, "routerRepo.Update.QueryRowxContext")
	}

	return comm, nil
}

// Delete comment
func (r *routerRepo) Delete(ctx context.Context, commentID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "routerRepo.Delete")
	defer span.Finish()

	result, err := r.db.ExecContext(ctx, deleteComment, commentID)
	if err != nil {
		return errors.Wrap(err, "routerRepo.Delete.ExecContext")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "routerRepo.Delete.RowsAffected")
	}

	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "routerRepo.Delete.rowsAffected")
	}

	return nil
}

// GetByID comment
func (r *routerRepo) GetByID(ctx context.Context, commentID uuid.UUID) (*models.CommentBase, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "routerRepo.GetByID")
	defer span.Finish()

	comment := &models.CommentBase{}
	if err := r.db.GetContext(ctx, comment, getCommentByID, commentID); err != nil {
		return nil, errors.Wrap(err, "routerRepo.GetByID.GetContext")
	}
	return comment, nil
}

// GetAllByNewsID comments
func (r *routerRepo) GetAllByNewsID(ctx context.Context, newsID uuid.UUID, query *utils.PaginationQuery) (*models.CommentsList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "routerRepo.GetAllByNewsID")
	defer span.Finish()

	var totalCount int
	if err := r.db.QueryRowContext(ctx, getTotalCountByNewsID, newsID).Scan(&totalCount); err != nil {
		return nil, errors.Wrap(err, "routerRepo.GetAllByNewsID.QueryRowContext")
	}
	if totalCount == 0 {
		return &models.CommentsList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
			Page:       query.GetPage(),
			Size:       query.GetSize(),
			HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
			Comments:   make([]*models.CommentBase, 0),
		}, nil
	}

	rows, err := r.db.QueryxContext(ctx, getCommentsByNewsID, newsID, query.GetOffset(), query.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "routerRepo.GetAllByNewsID.QueryxContext")
	}
	defer rows.Close()

	commentsList := make([]*models.CommentBase, 0, query.GetSize())
	for rows.Next() {
		comment := &models.CommentBase{}
		if err = rows.StructScan(comment); err != nil {
			return nil, errors.Wrap(err, "routerRepo.GetAllByNewsID.StructScan")
		}
		commentsList = append(commentsList, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "routerRepo.GetAllByNewsID.rows.Err")
	}

	return &models.CommentsList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),
		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
		Comments:   commentsList,
	}, nil
}
