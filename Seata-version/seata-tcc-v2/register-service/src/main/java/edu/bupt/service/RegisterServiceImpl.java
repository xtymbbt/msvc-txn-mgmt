package edu.bupt.service;

import edu.bupt.domain.Register;
import edu.bupt.feign.ProfileClient;
import edu.bupt.feign.EasyIdGeneratorClient;
import edu.bupt.feign.UserInfoClient;
import edu.bupt.tcc.RegisterTccAction;
import io.seata.spring.annotation.GlobalTransactional;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class RegisterServiceImpl implements RegisterService {
    // @Autowired
    // private RegisterMapper orderMapper;
    @Autowired
    EasyIdGeneratorClient easyIdGeneratorClient;
    @Autowired
    private ProfileClient profileClient;
    @Autowired
    private UserInfoClient userInfoClient;

    @Autowired
    private RegisterTccAction registerTccAction;

    @GlobalTransactional
    @Override
    public void register(Register register) {
        // 从全局唯一id发号器获得id
        Long userId = easyIdGeneratorClient.nextId("register_id");
        register.setId(userId);

        // orderMapper.create(register);

        // 这里修改成调用 TCC 第一节端方法
        registerTccAction.prepareCreateOrder(
                null,
                register.getId(),
                register.getUsername(),
                register.getPassword(),
                register.getPhoneNumber(),
                register.getEmail());


        // 创建用户信息user info
        userInfoClient.createUserInfo(register.getProductId(), register.getCount());

        // 创建用户档案user profile
        profileClient.decrease(register.getUserId(), register.getMoney());

    }
}