package callback

import (
	"github.com/gofiber/fiber/v2"
	"user/configs"
	"user/entities"
)

type Handlers struct {
	config *configs.AppConfig
	//transactionRepo    transaction.Repo
	//memberConnector    memberconnector.MemberConnector
	//walletConnector    walletconnector.WalletConnector
	//reportConnector    reportconnector.ReportConnector
	//router             map[string]func(c *fiber.Ctx, requestMessage []byte) error
	//db                 *gorm.DB
	//transactionService transactionService.Service
}

func New(
	config *configs.AppConfig,
	// transactionRepo transaction.Repo,
	// memberConnector memberconnector.MemberConnector,
	// walletConnector walletconnector.WalletConnector,
	// reportConnector reportconnector.ReportConnector,
	// db *gorm.DB,
	// transactionService transactionService.Service,
) *Handlers {
	handler := &Handlers{
		config: config,
		//transactionRepo:    transactionRepo,
		//memberConnector:    memberConnector,
		//walletConnector:    walletConnector,
		//reportConnector:    reportConnector,
		//db:                 db,
		//transactionService: transactionService,
	}
	//handler.router = map[string]func(c *fiber.Ctx, requestMessage []byte) error{
	//	//constant.ActionGetBalance:       handler.getBalance,
	//	//constant.ActionBet:              handler.bet,
	//	//constant.ActionSettle:           handler.settle,
	//	//constant.ActionUnsettle:         handler.unsettle,
	//	//constant.ActionResettle:         handler.resettle,
	//	//constant.ActionCancelBet:        handler.cancelBet,
	//	//constant.ActionVoidBet:          handler.voidBet,
	//	//constant.ActionVoidSettle:       handler.voidSettle,
	//	//constant.ActionTip:              handler.tip,
	//	//constant.ActionCancelTip:        handler.cancelTip,
	//	//constant.ActionGive:             handler.give,
	//	//constant.ActionBetNSettle:       handler.betNSettle,
	//	//constant.ActionCancelBetNSettle: handler.cancelBetNSettle,
	//	//constant.ActionFreeSpin:         handler.freeSpin,
	//}
	return handler
}

func (h *Handlers) response(c *fiber.Ctx, r *entities.Response) error {
	return c.Status(fiber.StatusOK).JSON(r)
}

func (h *Handlers) responseError(c *fiber.Ctx, err error) error {
	//switch expr := err.(type) {
	//case *transactionService.TxError:
	//	{
	//		errMsg := expr.Error()
	//		if errMsg == "" {
	//			return c.Status(fiber.StatusOK).JSON(new(entities.Response).Create(expr.ErrorCode, nil))
	//		}
	//		return c.Status(fiber.StatusOK).JSON(new(entities.Response).Create(expr.ErrorCode, expr))
	//	}
	//}
	return c.Status(fiber.StatusOK).JSON(new(entities.Response).Create(entities.Fail, err))
}
