// Package strategy 策略模式
package strategy

import (
	"fmt"
)

// FlyBehavior 飞行为的接口
type FlyBehavior interface {
	fly()
}

// QuackBehavior 叫行为的接口
type QuackBehavior interface {
	quack()
}

// FlyWithWings 为每个接口定义几种策略
type FlyWithWings struct {
}

// FlyNoWay 不会飞
type FlyNoWay struct {
}

// 每种struct实现FlyBehavior接口 代表了一种策略
func (fw *FlyWithWings) fly() {
	fmt.Println("利用翅膀飞翔")
}

func (fn *FlyNoWay) fly() {
	fmt.Println("不会飞")
}

// Quack 呱呱叫
type Quack struct {
}

// Squeak 吱吱叫
type Squeak struct {
}

// MuteQuack 不会叫
type MuteQuack struct {
}

// 实现QuackBehavior 呱呱叫的策略
func (q *Quack) quack() {
	fmt.Println("呱呱")
}

// 吱吱叫策略
func (s *Squeak) quack() {
	fmt.Println("吱吱")
}

func (mq *MuteQuack) quack() {
	fmt.Println("sorry 我不会叫")
}
