package presaleorder

import (
	"context"
	"fmt"
	"models/presaleorder"
	"systemlink/entities"
	"time"
)

func (s ServiceImpl) CreatePreSaleOrder(ctx context.Context, req presaleorder.PreSaleOrderRequest) (*presaleorder.PreSaleOrderRequest, error) {
	currentTime := time.Now()
	orderYear := currentTime.Format("20060102")
	orderNo := "MPOS-" + orderYear[2:6]
	poNo := currentTime.Format("150405") + "/" + currentTime.Format("2006")
	fmt.Println(poNo)
	preSaleNo, _ := s.presaleorderRepo.GetPreSaleOrderNo(ctx, orderNo)
	fmt.Println("preSaleNo = " + preSaleNo)
	switch len(preSaleNo) {
	case 1:
		preSaleNo = orderNo + "0000" + preSaleNo
	case 2:
		preSaleNo = orderNo + "000" + preSaleNo
	case 3:
		preSaleNo = orderNo + "00" + preSaleNo
	case 4:
		preSaleNo = orderNo + "0" + preSaleNo
	}
	fmt.Println("preSaleNo = " + preSaleNo)
	order := entities.PreSaleOrder{
		FTSaleOrderNo:   preSaleNo,
		FTInsUser:       req.FTInsUser,
		FDInsDate:       currentTime.Format("2006/01/02"),
		FTInsTime:       currentTime.Format("15:04:05"),
		FTQuotationNo:   "",
		FDSaleOrderDate: currentTime.Format("2006/01/02"),
		// wait for information
		// FTSaleOrderBy:   vUserinfo.FTEmpname,
		// FNMSysSaleId:    vUserinfo.FNMSysEmpID,
	}
	fmt.Println(order)
	return nil, nil
}
