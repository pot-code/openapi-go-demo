// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package flow

import flow "flow-editor-server/gen/flow"

type ConverterImpl struct{}

func (c *ConverterImpl) FlowSliceToFlowList(source []*Flow) []*flow.FlowListItemData {
	var pFlowFlowListItemDataList []*flow.FlowListItemData
	if source != nil {
		pFlowFlowListItemDataList = make([]*flow.FlowListItemData, len(source))
		for i := 0; i < len(source); i++ {
			pFlowFlowListItemDataList[i] = c.FlowToFlowListItem(source[i])
		}
	}
	return pFlowFlowListItemDataList
}
func (c *ConverterImpl) FlowToFlowDetail(source *Flow) *flow.FlowDetailData {
	var pFlowFlowDetailData *flow.FlowDetailData
	if source != nil {
		var flowFlowDetailData flow.FlowDetailData
		flowFlowDetailData.ID = UintToInt((*source).Model.ID)
		flowFlowDetailData.Title = (*source).Title
		var pString *string
		if (*source).Nodes != nil {
			xstring := *(*source).Nodes
			pString = &xstring
		}
		flowFlowDetailData.Nodes = pString
		var pString2 *string
		if (*source).Edges != nil {
			xstring2 := *(*source).Edges
			pString2 = &xstring2
		}
		flowFlowDetailData.Edges = pString2
		flowFlowDetailData.CreatedAt = TimeToString((*source).Model.CreatedAt)
		pFlowFlowDetailData = &flowFlowDetailData
	}
	return pFlowFlowDetailData
}
func (c *ConverterImpl) FlowToFlowListItem(source *Flow) *flow.FlowListItemData {
	var pFlowFlowListItemData *flow.FlowListItemData
	if source != nil {
		var flowFlowListItemData flow.FlowListItemData
		flowFlowListItemData.ID = UintToInt((*source).Model.ID)
		flowFlowListItemData.Title = (*source).Title
		flowFlowListItemData.CreatedAt = TimeToString((*source).Model.CreatedAt)
		pFlowFlowListItemData = &flowFlowListItemData
	}
	return pFlowFlowListItemData
}
