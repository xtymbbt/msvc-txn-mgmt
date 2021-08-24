package edu.bridge.mapper;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.UserInfo;

import java.util.HashMap;
import java.util.List;

public interface UserInfoMapper {
    boolean insertUserInfo(UserInfo userInfo,
                           CommonRequestBody commonRequestBody,
                           List<String> children);
}
