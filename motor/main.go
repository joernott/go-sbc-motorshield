/* Library for PiMotor Shield V2
   Developed by: Ott-Consult UG (haftungsbeschränkt)
   Author: Jörn Ott
   Project: RPi Motor Shield
   Copyright 2017 Joern Ott

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
package motor

import (
	"github.com/nathan-osman/go-rpigpio"
)

type pinset struct {
	e *rpi.Pin
	f *rpi.Pin
	r *rpi.Pin
}

type pinconfig struct {
	config [2]pinset
	arrow  *rpi.Pin
}

type motors map[string]pinconfig

type Motor struct {
	motorpins motors
	Test      bool
	Config    int
}

var Motors = [...]string{"MOTOR1", "MOTOR2", "MOTOR3", "MOTOR4"}

func newPinConfig(PowerPin int, ForwardPin int, BackPin int, ArrowPin int) (pinconfig, error) {
	var pc pinconfig
	var err error

	pc.config[0].e, err = rpi.OpenPin(PowerPin, rpi.OUT)
	if err != nil {
		return pc, err
	}
	pc.config[1].e = pc.config[0].e
	pc.config[0].f, err = rpi.OpenPin(ForwardPin, rpi.OUT)
	if err != nil {
		return pc, err
	}
	pc.config[1].r = pc.config[0].f
	pc.config[0].r, err = rpi.OpenPin(BackPin, rpi.OUT)
	if err != nil {
		return pc, err
	}
	pc.config[1].f = pc.config[0].r
	pc.arrow, err = rpi.OpenPin(ArrowPin, rpi.OUT)
	if err != nil {
		return pc, err
	}
	return pc, nil
}

func NewMotor() (Motor, error) {
	var m Motor
	var err error
	var pc pinconfig

	m.motorpins = make(motors)
	pc, err = newPinConfig(17, 22, 27, 16) //physical pins 11, 15, 13 and 36
	if err != nil {
		return m, err
	}
	m.motorpins["MOTOR1"] = pc
	pc, err = newPinConfig(25, 23, 24, 26) //physical pins 22, 16, 18 and 37
	if err != nil {
		return m, err
	}
	m.motorpins["MOTOR2"] = pc
	pc, err = newPinConfig(10, 9, 11, 19) //physical pins 19, 21, 23 and 35
	if err != nil {
		return m, err
	}
	m.motorpins["MOTOR3"] = pc

	pc, err = newPinConfig(12, 8, 7, 13) //physical pins 32, 24, 26 and 33
	if err != nil {
		return m, err
	}
	m.motorpins["MOTOR4"] = pc
	m.Config = 0
	return m, nil
}

func (m Motor) CloseMotor() {

	m.motorpins["MOTOR1"].config[0].e.Close()
	m.motorpins["MOTOR1"].config[0].f.Close()
	m.motorpins["MOTOR1"].config[0].r.Close()
	m.motorpins["MOTOR1"].arrow.Close()
}

func (m Motor) Forward(MotorName string) {
	m.motorpins[MotorName].config[m.Config].e.Write(rpi.HIGH)
	m.motorpins[MotorName].config[m.Config].f.Write(rpi.HIGH)
	m.motorpins[MotorName].config[m.Config].r.Write(rpi.LOW)
}

func (m Motor) Reverse(MotorName string) {
	m.motorpins[MotorName].config[m.Config].e.Write(rpi.HIGH)
	m.motorpins[MotorName].config[m.Config].f.Write(rpi.LOW)
	m.motorpins[MotorName].config[m.Config].r.Write(rpi.HIGH)
}

func (m Motor) Stop(MotorName string) {
	m.motorpins[MotorName].config[m.Config].e.Write(rpi.LOW)
	m.motorpins[MotorName].config[m.Config].f.Write(rpi.LOW)
	m.motorpins[MotorName].config[m.Config].r.Write(rpi.LOW)
}

func (m Motor) ArrowOn(MotorName string) {
	m.motorpins[MotorName].arrow.Write(rpi.HIGH)
}

func (m Motor) ArrowOff(MotorName string) {
	m.motorpins[MotorName].arrow.Write(rpi.LOW)
}
