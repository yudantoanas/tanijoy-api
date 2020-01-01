package object

import (
	"apijoy/model"
	"github.com/graphql-go/graphql"
	"strconv"
)

var ProjectStatusType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ProjectStatus",
	Description: "Menyediakan data proyek tersedia, berjalan dan selesai",
	Fields: graphql.Fields{
		"project_id": &graphql.Field{
			Description: "ID dari project",
			Type:        graphql.Int,
		},
		"jumlah": &graphql.Field{
			Description: "jumlah project",
			Type:        graphql.Int,
		},
	},
})

var ProjectRiskType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ProjectRisk",
	Description: "Object ProjectRisk",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description: "ID dari risk",
			Type:        graphql.Int,
		},
		"risk": &graphql.Field{
			Description: "nama risk",
			Type:        graphql.String,
		},
		"mitigation": &graphql.Field{
			Description: "mitigasi dari risk",
			Type:        graphql.String,
		},
	},
})

var ProjetDescriptionType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ProjectDescription",
	Description: "Object ProjectDescription",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description: "ID dari description",
			Type:        graphql.Int,
		},
		"label": &graphql.Field{
			Description: "label description",
			Type:        graphql.String,
		},
		"content": &graphql.Field{
			Description: "content description",
			Type:        graphql.String,
		},
		"order": &graphql.Field{
			Description: "order description",
			Type:        graphql.Int,
		},
	},
})

var ProjectFilesType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ProjectFiles",
	Description: "Object ProjectFiles",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description: "ID dari project files",
			Type:        graphql.Int,
		},
		"uploaded_by": &graphql.Field{
			Description: "Uploaded by user",
			Type:        UserType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				userId := p.Source.(model.Project_files).Uploaded_by

				var user model.User
				Database.First(&user, userId)

				return user, nil
			},
		},
		"filename": &graphql.Field{
			Description: "nama file",
			Type:        graphql.String,
		},
		"path": &graphql.Field{
			Description: "path file",
			Type:        graphql.String,
		},
		"label": &graphql.Field{
			Description: "label file",
			Type:        graphql.String,
		},
		"purpose": &graphql.Field{
			Description: "file purpose",
			Type:        graphql.String,
		},
	},
})

var ProjectUpdateType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "project update",
	Description: "update project in project details",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description: "ID dari Project_id",
			Type:        graphql.Int,
		},
		"category": &graphql.Field{
			Description: "kategori",
			Type:        graphql.String,
		},
		"additional_information": &graphql.Field{
			Description: "informasi tambahan",
			Type:        graphql.String,
		},
	},
})

var ProjectType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Project",
	Description: "Object Project_id untuk menyimpan data project tanijoy",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description: "ID dari Project_id",
			Type:        graphql.Int,
		},
		"name": &graphql.Field{
			Description: "Nama dari Project_id",
			Type:        graphql.String,
		},
		"amount": &graphql.Field{
			Description: "Jumlah dari Project_id",
			Type:        graphql.Int,
		},
		"risk": &graphql.Field{
			Description: "risk level",
			Type:        graphql.String,
		},
		"label": &graphql.Field{
			Description: "label dari Project_id",
			Type:        graphql.String,
		},
		"order": &graphql.Field{
			Description: "order dari Project_id",
			Type:        graphql.String,
		},
		"return_amount": &graphql.Field{
			Description: "Ekspektasi return dalam rupiah",
			Type:        graphql.Float,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				amount := p.Source.(model.Project).Amount
				expect := p.Source.(model.Project).Return_expected

				expectedAmmount := ((float64(expect) / 100) * float64(amount)) + float64(amount)

				return expectedAmmount, nil
			},
		},
		"return_expected": &graphql.Field{
			Description: "Ekspektasi return dari Project_id",
			Type:        graphql.Int,
		},
		"return_min": &graphql.Field{
			Description: "Minimal return dari Project_id",
			Type:        graphql.Int,
		},
		"return_max": &graphql.Field{
			Description: "Maximal return dari Project_id",
			Type:        graphql.Int,
		},
		"area": &graphql.Field{
			Description: "Area dari Project_id",
			Type:        graphql.Int,
		},
		"period": &graphql.Field{
			Description: "Periode dari Project_id",
			Type:        graphql.Int,
		},
		"planting_system": &graphql.Field{
			Description: "Sistem penanaman dari Project_id",
			Type:        graphql.String,
		},
		"plants_count": &graphql.Field{
			Description: "Jumlah penanaman dari Project_id",
			Type:        graphql.Int,
		},
		"image_path": &graphql.Field{
			Description: "Image path dari Project_id",
			Type:        graphql.String,
		},
		"thumbnail_path": &graphql.Field{
			Description: "Thumbnail path dari Project_id",
			Type:        graphql.String,
		},
		"commodity": &graphql.Field{
			Description: "Nama komoditas dari Project_id",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var commodity model.Commodity

				commId := p.Source.(model.Project).Commodities_id
				Database.First(&commodity, commId)

				return commodity.Name, nil
			},
		},
		"update_quantity": &graphql.Field{
			Description: "investment update proyek added/minus",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				projectId := p.Source.(model.Project).ID

				var count int
				Database.Raw("select sum(land_project_logs.qty) as quantity from investment_update " +
					"cross join land_project_logs " +
					"where project_id = $1 " +
					"and category = 'Added'", projectId).Row().Scan(&count)

				info := strconv.Itoa(count) + " Proyek Ditambah"

				return info, nil
			},
		},
		"update_selesai": &graphql.Field{
			Description: "investment update proyek selesai",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				projectId := p.Source.(model.Project).ID

				var count int
				Database.Raw("select count(*) from investment_update " +
					"where related_id in (select id from investments where project_id = $1) " +
					"and category = 'Finished'", projectId).Row().Scan(&count)

				info := strconv.Itoa(count) + " Proyek Selesai"

				return info, nil
			},
		},
		"update_terdanai": &graphql.Field{
			Description: "investment update proyek terdanai",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				projectId := p.Source.(model.Project).ID

				var count int
				Database.Raw("select count(id) from investment_update " +
					"where related_id in (select id from investment_concurrents where project_id = $1) " +
					"and category = 'Funded'", projectId).Row().Scan(&count)

				info := strconv.Itoa(count) + " Proyek Terdanai"

				return info, nil
			},
		},
		"descriptions": &graphql.Field{
			Description: "description dari Project_id",
			Type:        ProjetDescriptionType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				projectId := p.Source.(model.Project).ID

				var description model.Project_descriptions
				Database.Where(&model.Project_descriptions{Project_id: projectId}).First(&description)

				return description, nil
			},
		},
		"land": &graphql.Field{
			Description: "Lokasi tanah proyek",
			Type:        LandType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Source.(model.Project).ID

				var land_project model.Land_project
				var land model.Land
				Database.Where(&model.Land_project{Project_id: id}).First(&land_project)
				Database.Where(&model.Land{ID: land_project.Land_id}).First(&land)

				return land, nil
			},
		},
		"document": &graphql.Field{
			Description: "dokumen rincian proyek",
			Type:        ProjectFilesType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				projectId := p.Source.(model.Project).ID

				var doc model.Project_files
				Database.Where(&model.Project_files{Project_id: projectId, Label: "document"}).First(&doc)

				return doc, nil
			},
		},
		"map": &graphql.Field{
			Description: "path maps proyek",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				projectId := p.Source.(model.Project).ID

				var files model.Project_files
				Database.Where(&model.Project_files{Project_id: projectId, Label: "image", Filename:"maps-proyek"}).First(&files)

				return files.Path, nil
			},
		},
		"files": &graphql.Field{
			Description: "files proyek",
			Type:        graphql.NewList(ProjectFilesType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				projectId := p.Source.(model.Project).ID

				var files []model.Project_files
				Database.Where(&model.Project_files{Project_id: projectId, Label: "image"}).Not("filename", "maps-proyek").Find(&files)

				return files, nil
			},
		},
		"list_risk": &graphql.Field{
			Description: "resiko proyek",
			Type:        graphql.NewList(ProjectRiskType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				projectId := p.Source.(model.Project).ID

				var risks []model.Project_risks
				Database.Where(&model.Project_risks{Project_id: projectId}).Find(&risks)

				return risks, nil
			},
		},
		"farmers": &graphql.Field{
			Description: "Petani yang ada di proyek",
			Type:        graphql.NewList(FarmerType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				projectId := p.Source.(model.Project).ID
				rows, _ := Database.Raw("select farmer_id from farmer_commodity "+
					"where commodity_id = (select commodities_id from proposals where id = $1)", projectId).Rows() // (*sql.Rows, error)

				var farmerId int
				var ids []int
				var farmers []model.Farmer
				for rows.Next() {
					rows.Scan(&farmerId)
					ids = append(ids, farmerId)
				}
				Database.Where(ids).Find(&farmers)

				return farmers, nil
			},
		},
		"proyek_tersedia": &graphql.Field{
			Description: "Jumlah proyek yang tersedia",
			Type:        ProjectStatusType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				projectId := p.Source.(model.Project).ID

				var projectStatus model.ProjectStatus

				rows, _ := Database.Raw("select project_id, sum(ready) as tersedia from land_project where project_id = $1 group by project_id", projectId).Rows()

				for rows.Next() {
					rows.Scan(&projectStatus.Project_id, &projectStatus.Jumlah)
				}

				return projectStatus, nil
			},
		},
		"proyek_berjalan": &graphql.Field{
			Description: "Jumlah proyek yang berjalan",
			Type:        ProjectStatusType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var projectStatus model.ProjectStatus

				rows, _ := Database.Raw("select project_id, count(project_id) as berjalan from investments where actual_finish_at is null" +
					" and stage_id in (select id from stages where name <> 'Selesai')" +
					" group by project_id").Rows()

				for rows.Next() {
					rows.Scan(&projectStatus.Project_id, &projectStatus.Jumlah)
				}

				return projectStatus, nil
			},
		},
		"proyek_selesai": &graphql.Field{
			Description: "Jumlah proyek yang selesai",
			Type:        ProjectStatusType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var projectStatus model.ProjectStatus

				rows, _ := Database.Raw("select project_id, count(project_id) as selesai from investments where actual_finish_at is not null" +
					" and stage_id = (select id from stages where name = 'Selesai')" +
					" group by project_id").Rows()

				for rows.Next() {
					rows.Scan(&projectStatus.Project_id, &projectStatus.Jumlah)
				}

				return projectStatus, nil
			},
		},
	},
})
