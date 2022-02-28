package edu.bupt.service;

import edu.bupt.domain.Profile;
import edu.bupt.domain.Register;
import edu.bupt.domain.UserInfo;
import edu.bupt.feign.ProfileClient;
import edu.bupt.feign.EasyIdGeneratorClient;
import edu.bupt.feign.UserInfoClient;
import edu.bupt.tcc.RegisterTccAction;
import io.seata.spring.annotation.GlobalTransactional;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
@Slf4j
public class RegisterServiceImpl implements RegisterService {
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
        log.info("准备工作完成，开始获取全局唯一ID");
        // 从全局唯一id发号器获得id
        Long userId = easyIdGeneratorClient.nextId("register_id");
        log.info("获取全局唯一ID成功");
        register.setId(userId);

        // orderMapper.create(register);

        log.info("register.getId() is: {}", register.getId());
        log.info("begin tcc prepare");
        // 这里修改成调用 TCC 第一节端方法
        registerTccAction.prepareRegister(
                null,
                register.getId(),
                register.getUsername(),
                register.getPassword(),
                register.getPhoneNumber(),
                register.getEmail());
        log.info("tcc prepare success.");

        UserInfo userInfo = new UserInfo();
        userInfo.setId(register.getId());
        userInfo.setUsername(register.getUsername());
        userInfo.setEmail(register.getEmail());
        userInfo.setPhoneNumber(register.getPhoneNumber());

        // 创建用户信息user info
        userInfoClient.createUserInfo(userInfo);

        Profile profile = new Profile();
        profile.setId(register.getId());
        profile.setUsername(register.getUsername());

        // 创建用户档案user profile
        profileClient.createProfile(profile);

    }
}