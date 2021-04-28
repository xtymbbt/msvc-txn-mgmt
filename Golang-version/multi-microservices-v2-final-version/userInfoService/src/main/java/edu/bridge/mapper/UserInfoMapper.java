package edu.bridge.mapper;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.UserInfo;

import java.util.HashMap;

public interface UserInfoMapper {
    boolean insertUserInfo(UserInfo userInfo,
                           CommonRequestBody commonRequestBody,
                           HashMap<String, Boolean> children);
}
