package edu.bridge.service.impl;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.CommonResult;
import edu.bridge.domain.UserInfo;
import edu.bridge.mapper.UserInfoMapper;
import edu.bridge.service.UserInfoService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashMap;

@Service
public class UserInfoServiceImpl implements UserInfoService {
    @Autowired
    private UserInfoMapper userInfoMapper;

    @Override
    public CommonResult recordUserInfo(UserInfo userInfo,
                                       CommonRequestBody commonRequestBody) {
        // === Transaction codes ===
        HashMap<String, Boolean> children = new HashMap<>();
        if (commonRequestBody.getChild() != null && !commonRequestBody.getChild().equals("")) {
            children.put(commonRequestBody.getChild(), true);
        }
        // === Transaction codes ===

        return userInfoMapper.insertUserInfo(userInfo, commonRequestBody, children)
                ? new CommonResult(200, "insert userInfo success")
                : new CommonResult(500, "insert userInfo failed");
    }
}
