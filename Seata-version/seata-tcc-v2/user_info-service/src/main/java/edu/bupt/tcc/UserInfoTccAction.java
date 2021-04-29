package edu.bupt.tcc;

import edu.bupt.domain.UserInfo;
import io.seata.rm.tcc.api.BusinessActionContext;
import io.seata.rm.tcc.api.BusinessActionContextParameter;
import io.seata.rm.tcc.api.LocalTCC;
import io.seata.rm.tcc.api.TwoPhaseBusinessAction;

@LocalTCC
public interface UserInfoTccAction {

    @TwoPhaseBusinessAction(name = "userInfoTccAction", commitMethod = "commit", rollbackMethod = "rollback")
    boolean prepareCreateUserInfo(BusinessActionContext businessActionContext,
                                  @BusinessActionContextParameter(paramName = "userInfo")UserInfo userInfo);

    boolean commit(BusinessActionContext businessActionContext);

    boolean rollback(BusinessActionContext businessActionContext);

}
