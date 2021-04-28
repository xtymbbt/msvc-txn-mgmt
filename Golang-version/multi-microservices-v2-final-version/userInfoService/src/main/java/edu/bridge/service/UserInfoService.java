package edu.bridge.service;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.CommonResult;
import edu.bridge.domain.UserInfo;

public interface UserInfoService {
    CommonResult recordUserInfo(UserInfo userInfo,
                                CommonRequestBody commonRequestBody);
}
