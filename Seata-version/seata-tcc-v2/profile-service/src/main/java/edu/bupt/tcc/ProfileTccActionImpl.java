package edu.bupt.tcc;

import edu.bupt.domain.Profile;
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
    public boolean prepareDecreaseAccount(BusinessActionContext businessActionContext, Profile profile) {
        log.info("创建 profile 第一阶段，预留资源 - "+businessActionContext.getXid());

        profile.setStatus(0);
        profileMapper.insertProfile(profile);

        //事务成功，保存一个标识，供第二阶段进行判断
        ResultHolder.setResult(getClass(), businessActionContext.getXid(), "p");
        return true;
    }

    @Transactional
    @Override
    public boolean commit(BusinessActionContext businessActionContext) {
        log.info("创建 profile 第二阶段提交，修改 注册 状态1 - "+businessActionContext.getXid());

        // 防止幂等性，如果commit阶段重复执行则直接返回
        if (ResultHolder.getResult(getClass(), businessActionContext.getXid()) == null) {
            return true;
        }

        //Long profileId = (Long) businessActionContext.getActionContext("profileId");
        long profileId = Long.parseLong(businessActionContext.getActionContext("profileId").toString());
        Profile profile = new Profile();
        profile.setId(profileId);
        profile.setStatus(1);
        profileMapper.updateProfile(profile);

        //提交成功是删除标识
        ResultHolder.removeResult(getClass(), businessActionContext.getXid());
        return true;
    }

    @Transactional
    @Override
    public boolean rollback(BusinessActionContext businessActionContext) {
        log.info("创建 order 第二阶段回滚，删除订单 - "+businessActionContext.getXid());

        //第一阶段没有完成的情况下，不必执行回滚
        //因为第一阶段有本地事务，事务失败时已经进行了回滚。
        //如果这里第一阶段成功，而其他全局事务参与者失败，这里会执行回滚
        //幂等性控制：如果重复执行回滚则直接返回
        if (ResultHolder.getResult(getClass(), businessActionContext.getXid()) == null) {
            return true;
        }

        //Long profileId = (Long) businessActionContext.getActionContext("profileId");
        long profileId = Long.parseLong(businessActionContext.getActionContext("profileId").toString());
        profileMapper.deleteProfileById(profileId);

        //回滚结束时，删除标识
        ResultHolder.removeResult(getClass(), businessActionContext.getXid());
        return true;
    }
}
