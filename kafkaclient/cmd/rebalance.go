package cmd

import (
	"context"
	"fmt"

	"awesome-go/kafkaclient/console"
	"awesome-go/kafkaclient/kafka"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/kmsg"
)

/**
 * @Author: LFM
 * @Date: 2023/9/24 16:59
 * @Since: 1.0.0
 * @Desc: TODO
 */

func Rebalanced(client *kgo.Client, topics []kmsg.AlterPartitionAssignmentsRequestTopic) {
	newService, err := kafka.NewService(client)

	if err != nil {
		fmt.Errorf("aaaa:%v", err.Error())
	}

	service, err2 := console.NewService(newService)

	if err2 != nil {
		fmt.Errorf("bbbb:%v", err2.Error())
	}

	reassignments, err3 := service.ListPartitionReassignments(context.TODO())

	if err3 != nil {
		fmt.Errorf("cccc:%v", err3.Error())
	}

	fmt.Println(reassignments)

	assignments, err4 := service.AlterPartitionAssignments(context.TODO(), topics)
	if err4 != nil {
		fmt.Errorf("dddd:%v", err4.Error())
	}
	// 检查分区重新分配的结果
	fmt.Println(assignments)

}
