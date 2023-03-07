package graphql

import (
	"context"
	"github.com/BlackRule/App-and-its-features-CRUD/common"
	"github.com/BlackRule/App-and-its-features-CRUD/entities/app_feature"
	"github.com/BlackRule/App-and-its-features-CRUD/graphql/gen"
)

type featureResolver struct{ *Resolver }

func (r *featureResolver) Apps(ctx context.Context, obj *app_feature.Feature) ([]app_feature.App, error) {
	var apps []app_feature.App
	//TODO handle errors after each step (
	result1 := r.db.Model(&obj)
	result2 := result1.Association("Apps")
	result3 := result2.Find(&apps)
	return apps, result3.Error
}

func (r *Resolver) Feature() gen.FeatureResolver {
	return &featureResolver{r}
}

func (r *mutationResolver) CreateFeature(ctx context.Context, name string) (*app_feature.Feature, error) {
	//TODO
	_, err := common.Require_auth(ctx)
	if err != nil {
		return nil, err
	}
	feature := &app_feature.Feature{
		Name: name,
	}
	err = r.db.Create(feature).Error
	if err != nil {
		return nil, err
	}
	return feature, err
}

func (r *queryResolver) Features(ctx context.Context) ([]app_feature.Feature, error) {
	var features []app_feature.Feature
	result := r.db.Find(&features)
	return features, result.Error
}

func (r *mutationResolver) UpdateFeature(ctx context.Context, id int, name string) (*app_feature.Feature, error) {
	//TODO
	_, _ = common.Require_auth(ctx)
	feature := app_feature.Feature{
		ID:   uint(id),
		Name: name,
	}
	return &feature, r.db.Save(feature).Error
}

func (r *mutationResolver) DeleteFeature(ctx context.Context, id int) (int, error) {
	//TODO
	_, _ = common.Require_auth(ctx)
	feature := app_feature.Feature{
		ID: uint(id),
	}
	return id, r.db.Delete(&feature).Error
}
