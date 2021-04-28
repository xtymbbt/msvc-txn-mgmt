package edu.bupt.tcc;

import edu.bupt.mapper.ProfileMapper;
import io.seata.rm.tcc.api.BusinessActionContext;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import java.math.BigDecimal;

@Component
@Slf4j
public class ProfileTccActionImpl implements ProfileTccAction {
    @Autowired
    private ProfileMapper profileMapper;

    @Transactional
    @Override
    public boolean prepareDecreaseAccount(BusinessActionContext businessActionContext, Long userId, BigDecimal money) {
        log.info("减少账户金额，第一阶段锁定金额，userId="+userId+"， money="+money);

        edu.bupt.domain.Profile profile = profileMapper.selectById(userId);
        if (profile.getResidue().compareTo(money) < 0) {
            throw new RuntimeException("账户金额不足");
        }

        /*
        余额-money
        冻结+money
         */
        profileMapper.updateFrozen(userId, profile.getResidue().subtract(money), profile.getFrozen().add(money));

        if (Math.random() < 0.999) {
//            throw new RuntimeException("模拟异常");
        }

        //保存标识
        ResultHolder.setResult(getClass(), businessActionContext.getXid(), "p");
        return true;
    }

    @Transactional
    @Override
    public boolean commit(BusinessActionContext businessActionContext) {

        long userId = Long.parseLong(businessActionContext.getActionContext("userId").toString());
        BigDecimal money =  new BigDecimal(businessActionContext.getActionContext("money").toString());
        log.info("减少账户金额，第二阶段，提交，userId="+userId+"， money="+money);

        //防止重复提交
        if (ResultHolder.getResult(getClass(), businessActionContext.getXid()) == null) {
            return true;
        }

        profileMapper.updateFrozenToUsed(userId, money);

        //删除标识
        ResultHolder.removeResult(getClass(), businessActionContext.getXid());
        return true;
    }

    @Transactional
    @Override
    public boolean rollback(BusinessActionContext businessActionContext) {
        long userId = Long.parseLong(businessActionContext.getActionContext("userId").toString());
        BigDecimal money =  new BigDecimal(businessActionContext.getActionContext("money").toString());

        //防止重复提交
        if (ResultHolder.getResult(getClass(), businessActionContext.getXid()) == null) {
            return true;
        }

        log.info("减少账户金额，第二阶段，回滚，userId="+userId+"， money="+money);

        profileMapper.updateFrozenToResidue(userId, money);

        //删除标识
        ResultHolder.removeResult(getClass(), businessActionContext.getXid());
        return true;
    }
}
