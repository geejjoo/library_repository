package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/geejjoo/library_repository/internal/config"
	"github.com/geejjoo/library_repository/internal/db/adapter/dao"
	"github.com/geejjoo/library_repository/internal/db/tabler"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
	"reflect"
)

type SQLAdapter struct {
	db         *sqlx.DB
	sqlBuilder sq.StatementBuilderType
}

func NewSQLAdapter(db *sqlx.DB, dbConf config.DB) *SQLAdapter {
	return &SQLAdapter{
		db:         db,
		sqlBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (s *SQLAdapter) GetValues(pointers []interface{}, names []string) ([]interface{}, error) {
	var values []interface{}
	for i, pointer := range pointers {
		value := reflect.Indirect(reflect.ValueOf(pointer)).Interface()
		if names[i] == "id" {
			continue
		} else if reflect.ValueOf(value).Kind() == reflect.Struct {
			jsonValue, err := json.Marshal(value)
			if err != nil {
				return nil, err
			}
			values = append(values, jsonValue)
		} else if reflect.ValueOf(value).Kind() == reflect.Slice && reflect.ValueOf(value).Type().Elem().Kind() == reflect.Struct {
			marshal, err := json.Marshal(value)
			if err != nil {
				return nil, err
			}
			values = append(values, marshal)
		} else if photoUrls, ok := value.([]string); ok {
			values = append(values, pq.Array(photoUrls))
		} else {
			values = append(values, value)
		}
	}
	return values, nil
}

func (s *SQLAdapter) Create(ctx context.Context, entity tabler.Tabler) error {
	structInfo := dao.GetStructInfo(entity)
	queryBuilder := s.sqlBuilder.Insert(entity.TableName())
	for _, value := range structInfo.Fields {
		if value != "id" {
			queryBuilder = queryBuilder.Columns(value)
		}
	}
	values, err := s.GetValues(structInfo.Pointers, structInfo.Fields)
	if err != nil {
		return err
	}
	queryBuilder = queryBuilder.Values(values...)
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = s.db.ExecContext(ctx, query, args...)
	return err
}

func (s *SQLAdapter) Update(ctx context.Context, entity tabler.Tabler, condition dao.Condition) error {
	structInfo := dao.GetStructInfo(entity)
	queryBuilder := s.sqlBuilder.Update(entity.TableName())
	values, err := s.GetValues(structInfo.Pointers, structInfo.Fields)
	if err != nil {
		return err
	}
	for i, field := range structInfo.Fields {
		if field != "id" {
			value := values[i-1]
			queryBuilder = queryBuilder.Set(field, value)
		}
	}
	if condition.Equal != nil {
		for key, eq := range condition.Equal {
			queryBuilder = queryBuilder.Where(sq.Eq{key: eq})
		}
	}
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = s.db.ExecContext(ctx, query, args...)
	return err
}

func (s *SQLAdapter) BuildSelect(tableName string, condition dao.Condition, fields ...string) (string, []interface{}, error) {
	queryRaw := s.sqlBuilder.Select(fields...).From(tableName)
	if condition.Equal != nil {
		for key, eq := range condition.Equal {
			queryRaw = queryRaw.Where(sq.Eq{key: eq})
		}
	}

	if condition.NotEqual != nil {
		for key, notEq := range condition.NotEqual {
			queryRaw = queryRaw.Where(sq.NotEq{key: notEq})
		}
	}
	if condition.Order != nil {
		for _, order := range condition.Order {
			direction := "DESC"
			if order.Asc {
				direction = "ASC"
			}
			queryRaw = queryRaw.OrderBy(fmt.Sprintf("%s %s", order.Field, direction))
		}
	}

	if condition.LimitOffset != nil {
		if condition.LimitOffset.Limit > 0 {
			queryRaw = queryRaw.Limit(uint64(condition.LimitOffset.Limit))
		}
		if condition.LimitOffset.Offset > 0 {
			queryRaw = queryRaw.Limit(uint64(condition.LimitOffset.Offset))
		}
	}
	return queryRaw.ToSql()
}

func (s *SQLAdapter) List(ctx context.Context, dest interface{}, entity tabler.Tabler, condition dao.Condition) error {
	fields := dao.GetStructInfo(entity).Fields
	query, args, err := s.BuildSelect(entity.TableName(), condition, fields...)
	if err != nil {
		log.Println(err)
		return err
	}
	err = s.db.SelectContext(ctx, dest, query, args...)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// TODO доделать как будет время. Вместо List(), когда нужно забрать лишь один объект(100% будет случай,
// TODO когда при выборке найдётся несколько значений, но я заберу лишь одно - тогда, скорее всего, List() будет лучше)
//func (s *SQLAdapter) GetBy(ctx context.Context,dest interface{},entity tabler.Tabler,condition dao.Condition)error{
//	s.db.GetContext()
//	return nil
//}

// TODO Помечает запись, как удалённую путём модификации поля deleted_At. Свапнуть с методом ниже, если будет нужно
//func (s *SQLAdapter) Delete(ctx context.Context, entity tabler.Tabler, condition dao.Condition) error {
//	err := s.Update(ctx, entity, condition)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	return nil
//}

func (s *SQLAdapter) Delete(ctx context.Context, entity tabler.Tabler, condition dao.Condition) error {
	queryBuilder := s.sqlBuilder.Delete(entity.TableName())
	for key, eq := range condition.Equal {
		queryBuilder = queryBuilder.Where(sq.Eq{key: eq})
	}
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
