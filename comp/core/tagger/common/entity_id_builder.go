// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2024-present Datadog, Inc.

package common

import (
	"github.com/DataDog/datadog-agent/comp/core/tagger/types"
	workloadmeta "github.com/DataDog/datadog-agent/comp/core/workloadmeta/def"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)

// BuildTaggerEntityID builds tagger entity id based on workloadmeta entity id
func BuildTaggerEntityID(entityID workloadmeta.EntityID) types.EntityID {
	switch entityID.Kind {
	case workloadmeta.KindContainer:
		return types.NewEntityID(ContainerID, entityID.ID)
	case workloadmeta.KindKubernetesPod:
		return types.NewEntityID(KubernetesPodUID, entityID.ID)
	case workloadmeta.KindECSTask:
		return types.NewEntityID(ECSTask, entityID.ID)
	case workloadmeta.KindContainerImageMetadata:
		return types.NewEntityID(ContainerImageMetadata, entityID.ID)
	case workloadmeta.KindProcess:
		return types.NewEntityID(Process, entityID.ID)
	case workloadmeta.KindKubernetesDeployment:
		return types.NewEntityID(KubernetesDeployment, entityID.ID)
	case workloadmeta.KindHost:
		return types.NewEntityID(Host, entityID.ID)
	case workloadmeta.KindKubernetesMetadata:
		return types.NewEntityID(KubernetesMetadata, entityID.ID)
	default:
		log.Errorf("can't recognize entity %q with kind %q; trying %s://%s as tagger entity",
			entityID.ID, entityID.Kind, entityID.ID, entityID.Kind)
		return types.NewEntityID(types.EntityIDPrefix(entityID.Kind), entityID.ID)
	}
}
