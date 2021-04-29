package edu.bupt.tcc;

import edu.bupt.domain.Profile;
import io.seata.rm.tcc.api.BusinessActionContext;
import io.seata.rm.tcc.api.BusinessActionContextParameter;
import io.seata.rm.tcc.api.LocalTCC;
import io.seata.rm.tcc.api.TwoPhaseBusinessAction;

@LocalTCC
public interface ProfileTccAction {

    @TwoPhaseBusinessAction(name = "paymentTccAction", commitMethod = "commit", rollbackMethod = "rollback")
    boolean prepareDecreaseAccount(BusinessActionContext businessActionContext,
                                   @BusinessActionContextParameter(paramName = "profile")Profile profile);

    boolean commit(BusinessActionContext businessActionContext);

    boolean rollback(BusinessActionContext businessActionContext);

}
