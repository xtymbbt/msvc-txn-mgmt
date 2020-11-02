package edu.bupt.tcc;

import io.seata.rm.tcc.api.BusinessActionContext;
import io.seata.rm.tcc.api.BusinessActionContextParameter;
import io.seata.rm.tcc.api.LocalTCC;
import io.seata.rm.tcc.api.TwoPhaseBusinessAction;

import java.math.BigDecimal;

@LocalTCC
public interface PaymentTccAction {

    @TwoPhaseBusinessAction(name = "paymentTccAction", commitMethod = "commit", rollbackMethod = "rollback")
    boolean prepareDecreaseAccount(BusinessActionContext businessActionContext,
                                   @BusinessActionContextParameter(paramName = "userId") Long userId,
                                   @BusinessActionContextParameter(paramName = "money") BigDecimal money);

    boolean commit(BusinessActionContext businessActionContext);

    boolean rollback(BusinessActionContext businessActionContext);

}