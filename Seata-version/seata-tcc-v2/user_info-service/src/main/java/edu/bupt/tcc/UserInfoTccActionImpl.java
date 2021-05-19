package edu.bupt.tcc;

import com.alibaba.fastjson.JSON;
import edu.bupt.domain.UserInfo;
import edu.bupt.mapper.UserInfoMapper;
import io.seata.rm.tcc.api.BusinessActionContext;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

@Component
@Slf4j
public class UserInfoTccActionImpl implements UserInfoTccAction {

    @Autowired
    private UserInfoMapper userInfoMapper;

    @Transactional
    @Override
    public boolean prepareCreateUserInfo(BusinessActionContext businessActionContext, UserInfo userInfo) {
        System.out.println(businessActionContext);
        log.info("创建用户信息，第一阶段，预留资源 - "+businessActionContext.getXid());

        userInfo.setStatus(0);
        userInfoMapper.insertUserInfo(userInfo);
       
        //保存标识
        ResultHolder.setResult(getClass(), businessActionContext.getXid(), "p");
        return true;
    }

    @Transactional
    @Override
    public boolean commit(BusinessActionContext businessActionContext) {
        log.info("创建 userInfo 第二阶段提交，修改 userInfo 状态1 - "+businessActionContext.getXid());

        // 防止幂等性，如果commit阶段重复执行则直接返回
        if (ResultHolder.getResult(getClass(), businessActionContext.getXid()) == null) {
            return true;
        }

        Object o = businessActionContext.getActionContext("userInfo");
//        System.out.println(o.toString());
//        System.out.println(o.getClass());
        UserInfo userInfo = JSON.toJavaObject((JSON) o, UserInfo.class);
        //Long userInfoId = (Long) businessActionContext.getActionContext("userInfoId");
        userInfo.setStatus(1);
        userInfoMapper.updateUserInfo(userInfo);

        //提交成功是删除标识
        ResultHolder.removeResult(getClass(), businessActionContext.getXid());
        return true;
    }

    @Transactional
    @Override
    public boolean rollback(BusinessActionContext businessActionContext) {
        log.info("创建 userInfo 第二阶段回滚，删除用户信息 - "+businessActionContext.getXid());

        //第一阶段没有完成的情况下，不必执行回滚
        //因为第一阶段有本地事务，事务失败时已经进行了回滚。
        //如果这里第一阶段成功，而其他全局事务参与者失败，这里会执行回滚
        //幂等性控制：如果重复执行回滚则直接返回
        if (ResultHolder.getResult(getClass(), businessActionContext.getXid()) == null) {
            return true;
        }

        Object o = businessActionContext.getActionContext("userInfo");
//        System.out.println(o.toString());
//        System.out.println(o.getClass());
        UserInfo userInfo = JSON.toJavaObject((JSON) o, UserInfo.class);
//        System.out.println(userInfo.getId());
        //Long userInfoId = (Long) businessActionContext.getActionContext("userInfoId");
        userInfoMapper.deleteUserInfoById(userInfo.getId());

        //回滚结束时，删除标识
        ResultHolder.removeResult(getClass(), businessActionContext.getXid());
        return true;
    }
}
