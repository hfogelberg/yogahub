package main

import (
	"fmt"
)

func (connection *Connection) createPose(pose *Pose) (e error) {
	fmt.Println("Attempting to save pose")
	err := connection.Db.C("poses").Insert(&pose)
	if err != nil {
		fmt.Printf("Failed to insert pose %s\n", err.Error())
		return err
	}
	fmt.Println("Pose saved to db")
	return nil
}

func (connection *Connection) getPoses() (p []*Pose, e error) {
	var poses []*Pose
	err := connection.Db.C("poses").Find(nil).All(&poses)
	if err != nil {
		fmt.Printf("Error getting poses %s\n", err.Error())
		return nil, err
	}

	return poses, nil
}
