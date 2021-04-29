package edu.bupt.service;

import edu.bupt.domain.UserInfo;
import edu.bupt.feign.EasyIdGeneratorClient;
import edu.bupt.tcc.UserInfoTccAction;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class UserInfoServiceImpl implements UserInfoService {
    @Autowired
    EasyIdGeneratorClient easyIdGeneratorClient;

    @Autowired
    private UserInfoTccAction userInfoTccAction;

    @Override
    public void create(UserInfo userInfo) throws Exception {
        // 从全局唯一id发号器获得id
        Long userId = easyIdGeneratorClient.nextId("user_info_id");
        userInfo.setId(userId);

        userInfoTccAction.prepareCreateUserInfo(null, userInfo);
    }

}
