package util

import (
    "errors"
    "example/Project3/internal/constant"
    "example/Project3/internal/dto"
    "example/Project3/internal/model"
    "fmt"

    // "strconv"
    "strings"

    "gorm.io/gorm"
)

func ParseFilters(filters string) (filterMap dto.SearchFilters, err error) {
    var i = 0
    filterMap = make(dto.SearchFilters)
    
    var filterArr []string
    // var filterArray []string
    if strings.Contains(filters, constant.CommaDelimiter) {
        filterArr = strings.Split(filters, constant.CommaDelimiter)
    } else{
        filterArr = strings.Split(filters, constant.AndOperatorDelimiter)
    }
    
    length := len(filterArr)
    if length>1 {
    for i = 1; i < length; i++ {
        if strings.Contains(filterArr[i], constant.SearchOperatorDelimiter) {
            filterArr = strings.Split(filters, constant.CommaDelimiter)
        } else if strings.Contains(filters, "]"){
            filterArr = strings.Split(filters, "],")
        } else{
			filterArr = strings.Split(filters, constant.AndOperatorDelimiter)
		}
    for _, filter := range filterArr {
        if filter == "" {
            continue
        }

        var columnAndOperator,value string
        if strings.Contains(filter, constant.KeyValueDelimiter) {
                splitFilter := strings.SplitN(filter, constant.KeyValueDelimiter, 2)
                columnAndOperator = splitFilter[0]
                value = splitFilter[1]
				if strings.Contains(value, "[") {
                	value = strings.Trim(value, "[]")
				}
        } else {
            columnAndOperator = filter
            value = ""
        }

        if !strings.Contains(columnAndOperator, constant.SearchOperatorDelimiter) {
            return nil, errors.New("invalid Search Operator")
        }

        column, operator := SplitStringFromBack(columnAndOperator, constant.SearchOperatorDelimiter)
        table := constant.Root

        if _, ok := filterMap[table]; !ok {
            filterMap[table] = make(map[string]map[string]string)
        }

        if _, ok := filterMap[table][operator]; !ok {
            filterMap[table][operator] = make(map[string]string)
        }
        filterMap[table][operator][column] = value
        i++
    }
}
} else{
    for _, filter := range filterArr {
        if filter == "" {
            continue
        }

        var columnAndOperator,value string
        if strings.Contains(filter, constant.KeyValueDelimiter) {
                splitFilter := strings.SplitN(filter, constant.KeyValueDelimiter, 2)
                columnAndOperator = splitFilter[0]
                value = splitFilter[1]
				if strings.Contains(value, "[") {
                	value = strings.Trim(value, "[]")
				}
        } else {
            columnAndOperator = filter
            value = ""
        }

        if !strings.Contains(columnAndOperator, constant.SearchOperatorDelimiter) {
            return nil, errors.New("invalid Search Operator")
        }

        column, operator := SplitStringFromBack(columnAndOperator, constant.SearchOperatorDelimiter)
        table := constant.Root

        if _, ok := filterMap[table]; !ok {
            filterMap[table] = make(map[string]map[string]string)
        }

        if _, ok := filterMap[table][operator]; !ok {
            filterMap[table][operator] = make(map[string]string)
        }
        filterMap[table][operator][column] = value
    }
}
    
    return filterMap, nil
}


func getWherePredicates(operator string, column string, value string) ([]func(db *gorm.DB) *gorm.DB, error) {
    var predicates []func(db *gorm.DB) *gorm.DB
    switch constant.SearchOperator(operator) {
    case constant.Equal:
        predicates = append(predicates, ColumnValEqual(column, value))
    case constant.NotEqual:
        predicates = append(predicates, ColumnValNotEqual(column, value))
    case constant.In:
        values := strings.Split(value, constant.CommaDelimiter)
        predicates = append(predicates, ColumnValIn(column, values, constant.QueryAnd))
    case constant.NotIn:
        values := strings.Split(value, constant.CommaDelimiter)
        predicates = append(predicates, ColumnValNotIn(column, values, constant.QueryNot))
    case constant.GreaterThan:
        predicates = append(predicates, ColumnValGreaterThan(column, value, false))
    case constant.LessThan:
        predicates = append(predicates, ColumnValLessThan(column, value, false))
    case constant.GreaterThanEqualTo:
        predicates = append(predicates, ColumnValGreaterThan(column, value, true))
    case constant.LessThanEqualTo:
        predicates = append(predicates, ColumnValLessThan(column, value, true))
    case constant.Like:
        predicates = append(predicates, ColumnStrValStartsWith(column, value))
    case constant.IsNull:
        predicates = append(predicates, ColumnValNull(column))
    case constant.IsNotNull:
        predicates = append(predicates, ColumnValNotNull(column))
    default:
        return nil, errors.New("operator not defined")
    }
    return predicates, nil
}

func GetSortPredicates(sortParams string) []func(db *gorm.DB) *gorm.DB {
    var predicates []func(db *gorm.DB) *gorm.DB

    if sortParams == "" {
        sortParams = "id"
    }

    sortParamArray := strings.Split(sortParams, ",")
    for _, sortParam := range sortParamArray {
        order := constant.DefaultSortOrder
        column := sortParam
        if strings.Contains(sortParam, ".") {
            columnOrder := strings.Split(sortParam, ".")
            column = columnOrder[0]
            order = columnOrder[1]
        }
        predicates = append(predicates, ColumnOrderBy(column, order))
    }
    return predicates
}

func GetPaginationPredicates(page int, pageSize int) []func(db *gorm.DB) *gorm.DB {
    var predicates []func(db *gorm.DB) *gorm.DB
    if pageSize > 0 {
        predicates = append(predicates, Paginate(page, pageSize))
    }

    return predicates
}

func GetWherePredicates(query *gorm.DB, filterMap dto.SearchFilters, doc model.Document) (predicates []func(db *gorm.DB) *gorm.DB, err error) {
    for tableName, tableFilterMap := range filterMap {
        if tableName == constant.Root {
            tableName, err = GetTableName(query, doc)
            if err != nil {
                return nil, err
            }
        } else {
            joinCondition, err := ReadTag(doc, tableName, "join")
            if err != nil {
                return nil, err
            }
            query = query.Joins(joinCondition)
            tableName, err = ReadTag(doc, tableName, "tableName")
            if err != nil {
                return nil, err
            }
        }
        var operatorValue, colValue, value string
        for operator, colVal := range tableFilterMap {
            for col, val := range colVal {
                key := fmt.Sprintf("\"%s\".%s", tableName, col)
                // key := fmt.Sprintf("%s.%s", tableName, col)
                delete(colVal, col)
                colVal[key] = val
                operatorValue = operator
                colValue = col
                value = val
            }
        }

        tableWherePredicates, err := getWherePredicates(operatorValue, colValue, value)
        if err != nil {
            return nil, err
        }

        predicates = append(predicates, tableWherePredicates...)
    }
    return predicates, nil
}

func GetWherePredicatesUser(query *gorm.DB, filterMap dto.SearchFilters, user model.User) (predicates []func(db *gorm.DB) *gorm.DB, err error) {
    for tableName, tableFilterMap := range filterMap {
        if tableName == constant.Root {
            tableName, err = GetTableNameUser(query, user)
            if err != nil {
                return nil, err
            }
        } else {
            joinCondition, err := ReadTagUser(user, tableName, "join")
            if err != nil {
                return nil, err
            }
            query = query.Joins(joinCondition)
            tableName, err = ReadTagUser(user, tableName, "tableName")
            if err != nil {
                return nil, err
            }
        }

        for operator, colVal := range tableFilterMap {
            for col, val := range colVal {
                key := fmt.Sprintf("\"%s\".%s", tableName, col)
                // key := fmt.Sprintf("%s.%s", tableName, col)
                predicate, err := getWherePredicates(operator, key, val)
                if err != nil {
                    return nil, err
                }
                fmt.Println("check")
                predicates = append(predicates, predicate...)

            }
        }
    }

    return predicates, nil
}

func AddIncludes(includes string, query *gorm.DB) *gorm.DB {
    if includes == "" {
        return query
    }
    for _, include := range strings.Split(includes, ",") {
        query = query.Preload(include)
    }
    return query
}

func AddScopes(txn *gorm.DB, predicates ...[]func(db *gorm.DB) *gorm.DB) *gorm.DB {
    for _, predicate := range predicates {
        txn = txn.Scopes(predicate...)
    }
    return txn
}

