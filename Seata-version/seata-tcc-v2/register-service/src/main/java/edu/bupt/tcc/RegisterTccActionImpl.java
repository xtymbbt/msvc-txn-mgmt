package edu.bupt.tcc;

import edu.bupt.domain.Register;
import edu.bupt.mapper.RegisterMapper;
import io.seata.rm.tcc.api.BusinessActionContext;
import io.seata.rm.tcc.api.BusinessActionContextParameter;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import java.math.BigDecimal;

@Component
@Slf4j
public class RegisterTccActionImpl implements RegisterTccAction {
    @Autowired
    private RegisterMapper registerMapper;

    @Transactional
    @Override
    public boolean prepareRegister(BusinessActionContext businessActionContext,
                                   Long id,
                                   String username,
                                   String password,
                                   Long phoneNumber,
                                   String email) {
        log.info("创建 register 第一阶段，预留资源 - "+businessActionContext.getXid());

        Register register = new Register(id, username, password, phoneNumber, email, 0);
        registerMapper.insertRegister(register);

        //事务成功，保存一个标识，供第二阶段进行判断
        ResultHolder.setResult(getClass(), businessActionContext.getXid(), "p");
        return true;
    }

    @Transactional
    @Override
    public boolean commit(BusinessActionContext businessActionContext) {
        log.info("创建 register 第二阶段提交，修改 注册 状态1 - "+businessActionContext.getXid());

        // 防止幂等性，如果commit阶段重复执行则直接返回
        if (ResultHolder.getResult(getClass(), businessActionContext.getXid()) == null) {
            return true;
        }

        //Long registerId = (Long) businessActionContext.getActionContext("registerId");
        long registerId = Long.parseLong(businessActionContext.getActionContext("registerId").toString());
        Register register = new Register();
        register.setId(registerId);
        register.setStatus(1);
        registerMapper.updateRegister(register);

        //提交成功是删除标识
        ResultHolder.removeResult(getClass(), businessActionContext.getXid());
        return true;
    }

    @Transactional
    @Override
    public boolean rollback(BusinessActionContext businessActionContext) {
        log.info("创建 register 第二阶段回滚，删除订单 - "+businessActionContext.getXid());

        //第一阶段没有完成的情况下，不必执行回滚
        //因为第一阶段有本地事务，事务失败时已经进行了回滚。
        //如果这里第一阶段成功，而其他全局事务参与者失败，这里会执行回滚
        //幂等性控制：如果重复执行回滚则直接返回
        String result = ResultHolder.getResult(getClass(), businessActionContext.getXid());
        System.out.println(result);
        if (result == null) {
            return true;
        }

        //Long registerId = (Long) businessActionContext.getActionContext("id");
        long registerId = Long.parseLong(businessActionContext.getActionContext("id").toString());
        System.out.println("registerId: "+registerId);
        registerMapper.deleteRegisterById(registerId);
        System.out.println("rollback delete register successed.");

        //回滚结束时，删除标识
        ResultHolder.removeResult(getClass(), businessActionContext.getXid());
        return true;
    }
}
