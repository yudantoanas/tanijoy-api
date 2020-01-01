package queries

import (
	"apijoy/model"
	"apijoy/object"
	"github.com/graphql-go/graphql"
)

func GetAllProjectQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(object.ProjectType),
		Description: "Query untuk menampilkan list semua project",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var projects []model.Project
			Database.Find(&projects)

			return projects, nil
		},
	}
}

func ProjectDetail() *graphql.Field {
	return &graphql.Field{
		Type:        object.ProjectType,
		Description: "Query untuk menampilkan detail project",
		Args: graphql.FieldConfigArgument{
			"projectId": &graphql.ArgumentConfig{
				Type: graphql.Int,
				Description:"Parameter ID project",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			projectId := p.Args["projectId"].(int)

			var project model.Project
			Database.First(&project, projectId)

			return project, nil
		},
	}
}
