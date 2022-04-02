package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/satori/go.uuid"
)

type Appointment struct {
	Base
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	Recurring bool `json:"recurring"`
	StudentId uuid.UUID `json:"student_id" gorm:"type:uuid;column:student_id;not null"`
}

type AppointmentWrapper struct {
	Appointment
	Duration float64 `json:"duration"`
	StudentName string `json:"student_name"`
}

func (apt Appointment) GetCurrentWeekApt() (Appointment, error) {

	newLoc, err := time.LoadLocation("America/New_York"); if err != nil {
		return Appointment{}, err
	}
	
	defLoc, err := time.LoadLocation("UTC"); if err != nil {
		return Appointment{}, err
	}
	fmt.Println(apt.StartTime)
	var newStart time.Time
	var newEnd time.Time

	if !apt.Recurring {
		if (time.Now().After(apt.StartTime)){
			fmt.Println(time.Now(), apt.StartTime)
			return Appointment{}, errors.New("START TIME IN PAST")
		} 
		newStart = time.Date(apt.StartTime.Year(), apt.StartTime.Month(),  apt.StartTime.Day(), apt.StartTime.Hour(), apt.StartTime.Minute(), 0, 0, defLoc).In(newLoc)
		fmt.Println(newStart)
		newEnd = newStart.Add(apt.EndTime.Sub(apt.StartTime))
	} else {
		weekday := apt.StartTime.Weekday()
		hours, mins, _ := apt.StartTime.Clock()

		daysToAdd := ((weekday - time.Now().Weekday()) + 7) % 7

		newDay := time.Now().AddDate(0,0,int(daysToAdd))

		newStart = time.Date(newDay.Year(), newDay.Month(),  newDay.Day(), hours, mins, 0, 0, defLoc).In(newLoc)
		newEnd = newStart.Add(apt.EndTime.Sub(apt.StartTime))
	}

	curApt := Appointment{
		Base{apt.ID,
		time.Now()},
		newStart,
		newEnd,
		apt.Recurring,
		apt.StudentId,
	}

	fmt.Println(curApt)
	return curApt, nil
}
/*
func SendAppointmentReminder(currentApt Appointment) (bool, error){
	

}*/