package edu.bupt.tcc;

import com.alibaba.fastjson.JSON;
import edu.bupt.domain.Profile;
import edu.bupt.mapper.ProfileMapper;
import io.seata.rm.tcc.api.BusinessActionContext;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

@Component
@Slf4j
public class ProfileTccActionImpl implements ProfileTccAction {
    @Autowired
    private ProfileMapper profileMapper;

    @Transactional
    @Override
    public boolean prepareDecreaseAccount(BusinessActionContext businessActionContext, Profile profile) {
        log.info("创建 profile 第一阶段，预留资源 - "+businessActionContext.getXid());
        log.info("预留资源成功");

        profile.setStatus(0);
        log.info("一阶段写入数据库：");
        profileMapper.insertProfile(profile);
        log.info("一阶段写入数据库成功。");

        //事务成功，保存一个标识，供第二阶段进行判断
        ResultHolder.setResult(getClass(), businessActionContext.getXid(), "p");
        return true;
    }

    @Transactional
    @Override
    public boolean commit(BusinessActionContext businessActionContext) {
        log.info("创建 profile 第二阶段提交，修改 注册 状态1 - "+businessActionContext.getXid());
        log.info("二阶段提交，修改注册状态成功。");

        // 防止幂等性，如果commit阶段重复执行则直接返回
        if (ResultHolder.getResult(getClass(), businessActionContext.getXid()) == null) {
            return true;
        }

        Object o = businessActionContext.getActionContext("profile");
//        System.out.println(o.toString());
//        System.out.println(o.getClass());
        Profile profile  = JSON.toJavaObject((JSON) o, Profile.class);
        //Long profileId = (Long) businessActionContext.getActionContext("profileId");
        profile.setStatus(1);
        log.info("二阶段写入数据库");
        profileMapper.updateProfile(profile);
        log.info("二阶段写入数据库成功");

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

        Object o = businessActionContext.getActionContext("profile");
//        System.out.println(o.toString());
//        System.out.println(o.getClass());
        Profile profile  = JSON.toJavaObject((JSON) o, Profile.class);
        //Long profileId = (Long) businessActionContext.getActionContext("profile");
        profileMapper.deleteProfileById(profile.getId());

        //回滚结束时，删除标识
        ResultHolder.removeResult(getClass(), businessActionContext.getXid());
        return true;
    }
}
