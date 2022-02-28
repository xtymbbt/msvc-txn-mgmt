package edu.bupt.service;

import edu.bupt.domain.UserInfo;
import edu.bupt.feign.EasyIdGeneratorClient;
import edu.bupt.tcc.UserInfoTccAction;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
@Slf4j
public class UserInfoServiceImpl implements UserInfoService {
    @Autowired
    EasyIdGeneratorClient easyIdGeneratorClient;

    @Autowired
    private UserInfoTccAction userInfoTccAction;

    @Override
    public void create(UserInfo userInfo) throws Exception {
        log.info("tcc prepare begin");
        userInfoTccAction.prepareCreateUserInfo(null, userInfo);
        log.info("tcc prepare success");
    }

}
