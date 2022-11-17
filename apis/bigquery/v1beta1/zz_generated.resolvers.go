/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta11 "github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1"
	v1beta12 "github.com/upbound/provider-gcp/apis/kms/v1beta1"
	v1beta1 "github.com/upbound/provider-gcp/apis/sql/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Connection.
func (mg *Connection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.CloudSQL); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.CloudSQL[i3].Credential); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CloudSQL[i3].Credential[i4].Username),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.CloudSQL[i3].Credential[i4].UsernameRef,
				Selector:     mg.Spec.ForProvider.CloudSQL[i3].Credential[i4].UsernameSelector,
				To: reference.To{
					List:    &v1beta1.UserList{},
					Managed: &v1beta1.User{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.CloudSQL[i3].Credential[i4].Username")
			}
			mg.Spec.ForProvider.CloudSQL[i3].Credential[i4].Username = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.CloudSQL[i3].Credential[i4].UsernameRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.CloudSQL); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CloudSQL[i3].Database),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.CloudSQL[i3].DatabaseRef,
			Selector:     mg.Spec.ForProvider.CloudSQL[i3].DatabaseSelector,
			To: reference.To{
				List:    &v1beta1.DatabaseList{},
				Managed: &v1beta1.Database{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CloudSQL[i3].Database")
		}
		mg.Spec.ForProvider.CloudSQL[i3].Database = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CloudSQL[i3].DatabaseRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.CloudSQL); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CloudSQL[i3].InstanceID),
			Extract:      resource.ExtractParamPath("connection_name", true),
			Reference:    mg.Spec.ForProvider.CloudSQL[i3].InstanceIDRef,
			Selector:     mg.Spec.ForProvider.CloudSQL[i3].InstanceIDSelector,
			To: reference.To{
				List:    &v1beta1.DatabaseInstanceList{},
				Managed: &v1beta1.DatabaseInstance{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CloudSQL[i3].InstanceID")
		}
		mg.Spec.ForProvider.CloudSQL[i3].InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CloudSQL[i3].InstanceIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this DataTransferConfig.
func (mg *DataTransferConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DestinationDatasetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DestinationDatasetIDRef,
		Selector:     mg.Spec.ForProvider.DestinationDatasetIDSelector,
		To: reference.To{
			List:    &DatasetList{},
			Managed: &Dataset{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DestinationDatasetID")
	}
	mg.Spec.ForProvider.DestinationDatasetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DestinationDatasetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Dataset.
func (mg *Dataset) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Access); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Access[i3].Dataset); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Access[i3].Dataset[i4].Dataset); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Access[i3].Dataset[i4].Dataset[i5].DatasetID),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.Access[i3].Dataset[i4].Dataset[i5].DatasetIDRef,
					Selector:     mg.Spec.ForProvider.Access[i3].Dataset[i4].Dataset[i5].DatasetIDSelector,
					To: reference.To{
						List:    &DatasetList{},
						Managed: &Dataset{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.Access[i3].Dataset[i4].Dataset[i5].DatasetID")
				}
				mg.Spec.ForProvider.Access[i3].Dataset[i4].Dataset[i5].DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.Access[i3].Dataset[i4].Dataset[i5].DatasetIDRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Access); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Access[i3].Dataset); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Access[i3].Dataset[i4].Dataset); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Access[i3].Dataset[i4].Dataset[i5].ProjectID),
					Extract:      resource.ExtractParamPath("project", false),
					Reference:    mg.Spec.ForProvider.Access[i3].Dataset[i4].Dataset[i5].ProjectIDRef,
					Selector:     mg.Spec.ForProvider.Access[i3].Dataset[i4].Dataset[i5].ProjectIDSelector,
					To: reference.To{
						List:    &DatasetList{},
						Managed: &Dataset{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.Access[i3].Dataset[i4].Dataset[i5].ProjectID")
				}
				mg.Spec.ForProvider.Access[i3].Dataset[i4].Dataset[i5].ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.Access[i3].Dataset[i4].Dataset[i5].ProjectIDRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Access); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Access[i3].UserByEmail),
			Extract:      resource.ExtractParamPath("email", true),
			Reference:    mg.Spec.ForProvider.Access[i3].UserByEmailRef,
			Selector:     mg.Spec.ForProvider.Access[i3].UserByEmailSelector,
			To: reference.To{
				List:    &v1beta11.ServiceAccountList{},
				Managed: &v1beta11.ServiceAccount{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Access[i3].UserByEmail")
		}
		mg.Spec.ForProvider.Access[i3].UserByEmail = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Access[i3].UserByEmailRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.DefaultEncryptionConfiguration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DefaultEncryptionConfiguration[i3].KMSKeyName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DefaultEncryptionConfiguration[i3].KMSKeyNameRef,
			Selector:     mg.Spec.ForProvider.DefaultEncryptionConfiguration[i3].KMSKeyNameSelector,
			To: reference.To{
				List:    &v1beta12.CryptoKeyList{},
				Managed: &v1beta12.CryptoKey{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DefaultEncryptionConfiguration[i3].KMSKeyName")
		}
		mg.Spec.ForProvider.DefaultEncryptionConfiguration[i3].KMSKeyName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DefaultEncryptionConfiguration[i3].KMSKeyNameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this DatasetAccess.
func (mg *DatasetAccess) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Dataset); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Dataset[i3].Dataset); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Dataset[i3].Dataset[i4].DatasetID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.Dataset[i3].Dataset[i4].DatasetIDRef,
				Selector:     mg.Spec.ForProvider.Dataset[i3].Dataset[i4].DatasetIDSelector,
				To: reference.To{
					List:    &DatasetList{},
					Managed: &Dataset{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Dataset[i3].Dataset[i4].DatasetID")
			}
			mg.Spec.ForProvider.Dataset[i3].Dataset[i4].DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Dataset[i3].Dataset[i4].DatasetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Dataset); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Dataset[i3].Dataset); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Dataset[i3].Dataset[i4].ProjectID),
				Extract:      resource.ExtractParamPath("project", false),
				Reference:    mg.Spec.ForProvider.Dataset[i3].Dataset[i4].ProjectIDRef,
				Selector:     mg.Spec.ForProvider.Dataset[i3].Dataset[i4].ProjectIDSelector,
				To: reference.To{
					List:    &DatasetList{},
					Managed: &Dataset{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Dataset[i3].Dataset[i4].ProjectID")
			}
			mg.Spec.ForProvider.Dataset[i3].Dataset[i4].ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Dataset[i3].Dataset[i4].ProjectIDRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatasetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DatasetIDRef,
		Selector:     mg.Spec.ForProvider.DatasetIDSelector,
		To: reference.To{
			List:    &DatasetList{},
			Managed: &Dataset{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatasetID")
	}
	mg.Spec.ForProvider.DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatasetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserByEmail),
		Extract:      resource.ExtractParamPath("email", true),
		Reference:    mg.Spec.ForProvider.UserByEmailRef,
		Selector:     mg.Spec.ForProvider.UserByEmailSelector,
		To: reference.To{
			List:    &v1beta11.ServiceAccountList{},
			Managed: &v1beta11.ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserByEmail")
	}
	mg.Spec.ForProvider.UserByEmail = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserByEmailRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.View); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.View[i3].DatasetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.View[i3].DatasetIDRef,
			Selector:     mg.Spec.ForProvider.View[i3].DatasetIDSelector,
			To: reference.To{
				List:    &DatasetList{},
				Managed: &Dataset{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.View[i3].DatasetID")
		}
		mg.Spec.ForProvider.View[i3].DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.View[i3].DatasetIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.View); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.View[i3].ProjectID),
			Extract:      resource.ExtractParamPath("project", false),
			Reference:    mg.Spec.ForProvider.View[i3].ProjectIDRef,
			Selector:     mg.Spec.ForProvider.View[i3].ProjectIDSelector,
			To: reference.To{
				List:    &TableList{},
				Managed: &Table{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.View[i3].ProjectID")
		}
		mg.Spec.ForProvider.View[i3].ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.View[i3].ProjectIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.View); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.View[i3].TableID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.View[i3].TableIDRef,
			Selector:     mg.Spec.ForProvider.View[i3].TableIDSelector,
			To: reference.To{
				List:    &TableList{},
				Managed: &Table{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.View[i3].TableID")
		}
		mg.Spec.ForProvider.View[i3].TableID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.View[i3].TableIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this DatasetIAMBinding.
func (mg *DatasetIAMBinding) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatasetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DatasetIDRef,
		Selector:     mg.Spec.ForProvider.DatasetIDSelector,
		To: reference.To{
			List:    &DatasetList{},
			Managed: &Dataset{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatasetID")
	}
	mg.Spec.ForProvider.DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatasetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DatasetIAMMember.
func (mg *DatasetIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatasetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DatasetIDRef,
		Selector:     mg.Spec.ForProvider.DatasetIDSelector,
		To: reference.To{
			List:    &DatasetList{},
			Managed: &Dataset{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatasetID")
	}
	mg.Spec.ForProvider.DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatasetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DatasetIAMPolicy.
func (mg *DatasetIAMPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatasetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DatasetIDRef,
		Selector:     mg.Spec.ForProvider.DatasetIDSelector,
		To: reference.To{
			List:    &DatasetList{},
			Managed: &Dataset{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatasetID")
	}
	mg.Spec.ForProvider.DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatasetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Job.
func (mg *Job) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Copy); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Copy[i3].DestinationEncryptionConfiguration); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Copy[i3].DestinationEncryptionConfiguration[i4].KMSKeyName),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Copy[i3].DestinationEncryptionConfiguration[i4].KMSKeyNameRef,
				Selector:     mg.Spec.ForProvider.Copy[i3].DestinationEncryptionConfiguration[i4].KMSKeyNameSelector,
				To: reference.To{
					List:    &v1beta12.CryptoKeyList{},
					Managed: &v1beta12.CryptoKey{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Copy[i3].DestinationEncryptionConfiguration[i4].KMSKeyName")
			}
			mg.Spec.ForProvider.Copy[i3].DestinationEncryptionConfiguration[i4].KMSKeyName = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Copy[i3].DestinationEncryptionConfiguration[i4].KMSKeyNameRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Copy); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Copy[i3].DestinationTable); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Copy[i3].DestinationTable[i4].DatasetID),
				Extract:      resource.ExtractParamPath("dataset_id", false),
				Reference:    mg.Spec.ForProvider.Copy[i3].DestinationTable[i4].DatasetIDRef,
				Selector:     mg.Spec.ForProvider.Copy[i3].DestinationTable[i4].DatasetIDSelector,
				To: reference.To{
					List:    &TableList{},
					Managed: &Table{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Copy[i3].DestinationTable[i4].DatasetID")
			}
			mg.Spec.ForProvider.Copy[i3].DestinationTable[i4].DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Copy[i3].DestinationTable[i4].DatasetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Copy); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Copy[i3].DestinationTable); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Copy[i3].DestinationTable[i4].ProjectID),
				Extract:      resource.ExtractParamPath("project", false),
				Reference:    mg.Spec.ForProvider.Copy[i3].DestinationTable[i4].ProjectIDRef,
				Selector:     mg.Spec.ForProvider.Copy[i3].DestinationTable[i4].ProjectIDSelector,
				To: reference.To{
					List:    &TableList{},
					Managed: &Table{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Copy[i3].DestinationTable[i4].ProjectID")
			}
			mg.Spec.ForProvider.Copy[i3].DestinationTable[i4].ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Copy[i3].DestinationTable[i4].ProjectIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Copy); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Copy[i3].DestinationTable); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Copy[i3].DestinationTable[i4].TableID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.Copy[i3].DestinationTable[i4].TableIDRef,
				Selector:     mg.Spec.ForProvider.Copy[i3].DestinationTable[i4].TableIDSelector,
				To: reference.To{
					List:    &TableList{},
					Managed: &Table{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Copy[i3].DestinationTable[i4].TableID")
			}
			mg.Spec.ForProvider.Copy[i3].DestinationTable[i4].TableID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Copy[i3].DestinationTable[i4].TableIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Extract); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Extract[i3].SourceTable); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Extract[i3].SourceTable[i4].DatasetID),
				Extract:      resource.ExtractParamPath("dataset_id", false),
				Reference:    mg.Spec.ForProvider.Extract[i3].SourceTable[i4].DatasetIDRef,
				Selector:     mg.Spec.ForProvider.Extract[i3].SourceTable[i4].DatasetIDSelector,
				To: reference.To{
					List:    &TableList{},
					Managed: &Table{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Extract[i3].SourceTable[i4].DatasetID")
			}
			mg.Spec.ForProvider.Extract[i3].SourceTable[i4].DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Extract[i3].SourceTable[i4].DatasetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Extract); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Extract[i3].SourceTable); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Extract[i3].SourceTable[i4].ProjectID),
				Extract:      resource.ExtractParamPath("project", false),
				Reference:    mg.Spec.ForProvider.Extract[i3].SourceTable[i4].ProjectIDRef,
				Selector:     mg.Spec.ForProvider.Extract[i3].SourceTable[i4].ProjectIDSelector,
				To: reference.To{
					List:    &TableList{},
					Managed: &Table{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Extract[i3].SourceTable[i4].ProjectID")
			}
			mg.Spec.ForProvider.Extract[i3].SourceTable[i4].ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Extract[i3].SourceTable[i4].ProjectIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Extract); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Extract[i3].SourceTable); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Extract[i3].SourceTable[i4].TableID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.Extract[i3].SourceTable[i4].TableIDRef,
				Selector:     mg.Spec.ForProvider.Extract[i3].SourceTable[i4].TableIDSelector,
				To: reference.To{
					List:    &TableList{},
					Managed: &Table{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Extract[i3].SourceTable[i4].TableID")
			}
			mg.Spec.ForProvider.Extract[i3].SourceTable[i4].TableID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Extract[i3].SourceTable[i4].TableIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Load); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Load[i3].DestinationTable); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Load[i3].DestinationTable[i4].DatasetID),
				Extract:      resource.ExtractParamPath("dataset_id", false),
				Reference:    mg.Spec.ForProvider.Load[i3].DestinationTable[i4].DatasetIDRef,
				Selector:     mg.Spec.ForProvider.Load[i3].DestinationTable[i4].DatasetIDSelector,
				To: reference.To{
					List:    &TableList{},
					Managed: &Table{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Load[i3].DestinationTable[i4].DatasetID")
			}
			mg.Spec.ForProvider.Load[i3].DestinationTable[i4].DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Load[i3].DestinationTable[i4].DatasetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Load); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Load[i3].DestinationTable); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Load[i3].DestinationTable[i4].ProjectID),
				Extract:      resource.ExtractParamPath("project", false),
				Reference:    mg.Spec.ForProvider.Load[i3].DestinationTable[i4].ProjectIDRef,
				Selector:     mg.Spec.ForProvider.Load[i3].DestinationTable[i4].ProjectIDSelector,
				To: reference.To{
					List:    &TableList{},
					Managed: &Table{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Load[i3].DestinationTable[i4].ProjectID")
			}
			mg.Spec.ForProvider.Load[i3].DestinationTable[i4].ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Load[i3].DestinationTable[i4].ProjectIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Load); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Load[i3].DestinationTable); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Load[i3].DestinationTable[i4].TableID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.Load[i3].DestinationTable[i4].TableIDRef,
				Selector:     mg.Spec.ForProvider.Load[i3].DestinationTable[i4].TableIDSelector,
				To: reference.To{
					List:    &TableList{},
					Managed: &Table{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Load[i3].DestinationTable[i4].TableID")
			}
			mg.Spec.ForProvider.Load[i3].DestinationTable[i4].TableID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Load[i3].DestinationTable[i4].TableIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Query); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Query[i3].DefaultDataset); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Query[i3].DefaultDataset[i4].DatasetID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Query[i3].DefaultDataset[i4].DatasetIDRef,
				Selector:     mg.Spec.ForProvider.Query[i3].DefaultDataset[i4].DatasetIDSelector,
				To: reference.To{
					List:    &DatasetList{},
					Managed: &Dataset{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Query[i3].DefaultDataset[i4].DatasetID")
			}
			mg.Spec.ForProvider.Query[i3].DefaultDataset[i4].DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Query[i3].DefaultDataset[i4].DatasetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Query); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Query[i3].DestinationTable); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Query[i3].DestinationTable[i4].DatasetID),
				Extract:      resource.ExtractParamPath("dataset_id", false),
				Reference:    mg.Spec.ForProvider.Query[i3].DestinationTable[i4].DatasetIDRef,
				Selector:     mg.Spec.ForProvider.Query[i3].DestinationTable[i4].DatasetIDSelector,
				To: reference.To{
					List:    &TableList{},
					Managed: &Table{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Query[i3].DestinationTable[i4].DatasetID")
			}
			mg.Spec.ForProvider.Query[i3].DestinationTable[i4].DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Query[i3].DestinationTable[i4].DatasetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Query); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Query[i3].DestinationTable); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Query[i3].DestinationTable[i4].ProjectID),
				Extract:      resource.ExtractParamPath("project", false),
				Reference:    mg.Spec.ForProvider.Query[i3].DestinationTable[i4].ProjectIDRef,
				Selector:     mg.Spec.ForProvider.Query[i3].DestinationTable[i4].ProjectIDSelector,
				To: reference.To{
					List:    &TableList{},
					Managed: &Table{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Query[i3].DestinationTable[i4].ProjectID")
			}
			mg.Spec.ForProvider.Query[i3].DestinationTable[i4].ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Query[i3].DestinationTable[i4].ProjectIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Query); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Query[i3].DestinationTable); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Query[i3].DestinationTable[i4].TableID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.Query[i3].DestinationTable[i4].TableIDRef,
				Selector:     mg.Spec.ForProvider.Query[i3].DestinationTable[i4].TableIDSelector,
				To: reference.To{
					List:    &TableList{},
					Managed: &Table{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Query[i3].DestinationTable[i4].TableID")
			}
			mg.Spec.ForProvider.Query[i3].DestinationTable[i4].TableID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Query[i3].DestinationTable[i4].TableIDRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this ReservationAssignment.
func (mg *ReservationAssignment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Reservation),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ReservationRef,
		Selector:     mg.Spec.ForProvider.ReservationSelector,
		To: reference.To{
			List:    &ReservationList{},
			Managed: &Reservation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Reservation")
	}
	mg.Spec.ForProvider.Reservation = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ReservationRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Routine.
func (mg *Routine) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatasetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DatasetIDRef,
		Selector:     mg.Spec.ForProvider.DatasetIDSelector,
		To: reference.To{
			List:    &DatasetList{},
			Managed: &Dataset{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatasetID")
	}
	mg.Spec.ForProvider.DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatasetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Table.
func (mg *Table) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatasetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DatasetIDRef,
		Selector:     mg.Spec.ForProvider.DatasetIDSelector,
		To: reference.To{
			List:    &DatasetList{},
			Managed: &Dataset{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatasetID")
	}
	mg.Spec.ForProvider.DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatasetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TableIAMPolicy.
func (mg *TableIAMPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatasetID),
		Extract:      resource.ExtractParamPath("dataset_id", false),
		Reference:    mg.Spec.ForProvider.DatasetIDRef,
		Selector:     mg.Spec.ForProvider.DatasetIDSelector,
		To: reference.To{
			List:    &TableList{},
			Managed: &Table{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatasetID")
	}
	mg.Spec.ForProvider.DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatasetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Project),
		Extract:      resource.ExtractParamPath("project", false),
		Reference:    mg.Spec.ForProvider.ProjectRef,
		Selector:     mg.Spec.ForProvider.ProjectSelector,
		To: reference.To{
			List:    &TableList{},
			Managed: &Table{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Project")
	}
	mg.Spec.ForProvider.Project = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TableID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TableIDRef,
		Selector:     mg.Spec.ForProvider.TableIDSelector,
		To: reference.To{
			List:    &TableList{},
			Managed: &Table{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TableID")
	}
	mg.Spec.ForProvider.TableID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TableIDRef = rsp.ResolvedReference

	return nil
}
