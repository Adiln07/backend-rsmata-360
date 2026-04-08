package usecases

import (
	"backend-rsmata-360/models"
	"backend-rsmata-360/repositories"
	"backend-rsmata-360/responses"
)

func GetAllFloorsWithRooms() ([]responses.FloorWithRoomsChildrenResponse, error) {

	floors , err := repositories.MapGetAllFloors()
	if err != nil {
		return []responses.FloorWithRoomsChildrenResponse{}, err
	}

	floorRooms, err := repositories.MapGetFloorRoom()
	if err != nil {
		return []responses.FloorWithRoomsChildrenResponse{}, err
	}

	floorRoomMap := make(map[uint][]uint)
	for _, fr := range floorRooms{
		floorRoomMap[fr.FloorId] = append(floorRoomMap[fr.FloorId], fr.RoomId)
	}

	rooms, err := repositories.MapGetAllRooms()
	if err != nil {
		return []responses.FloorWithRoomsChildrenResponse{}, err
	}

	roomMap := make(map[uint]models.Room)
	roomIDs := []uint{}
	for _, room := range rooms{
		roomMap[room.Id] = room
		roomIDs = append(roomIDs, room.Id)
	}

	hotspotInformation, err := repositories.MapGetAllHotspotInformation(roomIDs)
	if err != nil {
		return []responses.FloorWithRoomsChildrenResponse{}, err
	}

	hotspotNavigation, err := repositories.MapGetAllHotspotNavigation(roomIDs)
	if err != nil {
		return []responses.FloorWithRoomsChildrenResponse{}, err
	}

	mapInfo := make(map[uint][]models.HotspotInformation)
	mapNav := make(map[uint][]models.HotspotNav)

	for _, info := range hotspotInformation{
		mapInfo[info.Room_Id] = append(mapInfo[info.Room_Id], info)
	}

	for _, nav := range hotspotNavigation{
		mapNav[nav.Room_Id] = append(mapNav[nav.Room_Id], nav)
	}

	// * response 
	var response []responses.FloorWithRoomsChildrenResponse

	for _, floor := range floors{
		var roomList []responses.RoomWithChildrenResponse
		roomIDs := floorRoomMap[floor.Id]

		for _, rid := range roomIDs{
			if room, ok := roomMap[rid]; ok{

				// *convert hotspot info -> response type
				var infoResponse []responses.HotspotInformationResponse
				for _, info := range mapInfo[room.Id]{
					infoResponse = append(infoResponse, responses.HotspotInformationResponse{
						Id: info.Id,
						Room_Id: info.Room_Id,
						Yaw: info.Yaw,
						Pitch: info.Pitch,
						Label: info.Label,
						Description: info.Description,
						Status: info.Status,
					})
				}

				// * convert hotspot nav -> response type
				var navResponse []responses.HotspotNavigationResponse
				for _, nav := range mapNav[room.Id]{
					navResponse = append(navResponse, responses.HotspotNavigationResponse{
						Id: nav.Id,
						Room_Id: nav.Room_Id,
						Yaw: nav.Yaw,
						Pitch: nav.Pitch,
						Description: nav.Description,
						Target_Room_Label: nav.Target_Room_Label,
						Target_Room_Id: nav.Target_Room_Id,
						Status: nav.Status,
					})
				}

				// * convert hotspot info & response type
				roomList = append(roomList, responses.RoomWithChildrenResponse{
					Id:room.Id,
					Name: room.Name,
					Image: room.Image,
					Pos_x: room.Pos_x,
					Pos_y: room.Pos_y,
					Status: room.Status,
					HotspotInfo: infoResponse,
					HotspotNav: navResponse,
				})
			}
		}

		response = append(response, responses.FloorWithRoomsChildrenResponse{
			Id: floor.Id,
			Name: floor.Name,
			FloorPlan: floor.FloorPlan,
			Status: floor.Status,
			Rooms: roomList,
		})
	}

	return response, nil
}

func GetFloorByIdWithRooms(floorId int)(responses.FloorWithRoomsChildrenResponse, error){

	floor, err := repositories.MapGetFloorById(floorId)
	if err != nil {
		return responses.FloorWithRoomsChildrenResponse{}, err
	}

	floorRooms, err := repositories.MapGetFloorRoomByFloorId(floorId)
	if err != nil {
		return responses.FloorWithRoomsChildrenResponse{}, err
	}

	roomIDs := []uint{}

	for _, fr := range floorRooms{
		roomIDs = append(roomIDs, fr.RoomId)
	}

	if len(roomIDs) == 0 {
	return responses.FloorWithRoomsChildrenResponse{
		Id: floor.Id,
		Name: floor.Name,
		FloorPlan: floor.FloorPlan,
		Status: floor.Status,
		Rooms: []responses.RoomWithChildrenResponse{},
	}, nil
}

	rooms, err := repositories.MapGetAllRoomById(roomIDs)

	if err != nil {
		return responses.FloorWithRoomsChildrenResponse{}, err
	}

	hotspotInformation, err := repositories.MapGetAllHotspotInformation(roomIDs)

	if err != nil {
		return responses.FloorWithRoomsChildrenResponse{}, err
	}

	hotspotNavigation, err := repositories.MapGetAllHotspotNavigation(roomIDs)

	if err != nil {
		return responses.FloorWithRoomsChildrenResponse{}, err
	}

	mapInfo := make(map[uint][]models.HotspotInformation)
	mapNav := make(map[uint][]models.HotspotNav)

	for _, info := range hotspotInformation{
		mapInfo[info.Room_Id] = append(mapInfo[info.Room_Id], info)
	}

	for _, nav := range hotspotNavigation{
		mapNav[nav.Room_Id] = append(mapNav[nav.Room_Id], nav)
	}

	var roomResponse []responses.RoomWithChildrenResponse
	
	for _, room := range rooms{

		infoList := mapInfo[room.Id]
		navList := mapNav[room.Id]

		var infoResponse []responses.HotspotInformationResponse
		for _, info := range infoList{
			infoResponse = append(infoResponse, responses.HotspotInformationResponse{
				Id: info.Id,
				Room_Id: info.Room_Id,
				Yaw : info.Yaw,
				Pitch: info.Pitch,
				Label : info.Label,
				Description: info.Description,
				Status: info.Status,
			})
		}

		var navResponse []responses.HotspotNavigationResponse
		for _, nav := range navList{
			navResponse = append(navResponse, responses.HotspotNavigationResponse{
				Id: nav.Id,
				Room_Id: nav.Room_Id,
				Yaw: nav.Yaw,
				Pitch:nav.Pitch,
				Description:nav.Description,
				Target_Room_Label: nav.Target_Room_Label,
				Target_Room_Id: nav.Target_Room_Id,
				Status: nav.Status,
			})
		}

		roomResponse = append(roomResponse, responses.RoomWithChildrenResponse{
			Id: room.Id,
			Name: room.Name,
			Image: room.Image,
			Pos_x: room.Pos_x,
			Pos_y: room.Pos_y,
			Status: room.Status,
			HotspotInfo: infoResponse,
			HotspotNav: navResponse,
		})
	}

	response := responses.FloorWithRoomsChildrenResponse{
		Id: floor.Id,
		Name: floor.Name,
		FloorPlan: floor.FloorPlan,
		Status: floor.Status,
		Rooms: roomResponse,
	}

	return response, nil
}

