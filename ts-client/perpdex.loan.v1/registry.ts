import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateParams } from "./types/perpdex/loan/v1/tx";
import { MsgRequestLoan } from "./types/perpdex/loan/v1/tx";
import { MsgCancelLoan } from "./types/perpdex/loan/v1/tx";
import { MsgApproveLoan } from "./types/perpdex/loan/v1/tx";
import { MsgRepayLoan } from "./types/perpdex/loan/v1/tx";
import { MsgLiquidateLoan } from "./types/perpdex/loan/v1/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/perpdex.loan.v1.MsgUpdateParams", MsgUpdateParams],
    ["/perpdex.loan.v1.MsgRequestLoan", MsgRequestLoan],
    ["/perpdex.loan.v1.MsgCancelLoan", MsgCancelLoan],
    ["/perpdex.loan.v1.MsgApproveLoan", MsgApproveLoan],
    ["/perpdex.loan.v1.MsgRepayLoan", MsgRepayLoan],
    ["/perpdex.loan.v1.MsgLiquidateLoan", MsgLiquidateLoan],
    
];

export { msgTypes }