package edu.bridge.service.impl;

import edu.bridge.domain.*;
import edu.bridge.feign.ProfileService;
import edu.bridge.feign.UserInfoService;
import edu.bridge.mapper.RegisterMapper;
import edu.bridge.service.RegisterService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;

@Service
public class RegisterServiceImpl implements RegisterService {
    @Autowired
    private RegisterMapper registerMapper;
    @Autowired
    private UserInfoService userInfoService;
    @Autowired
    private ProfileService profileService;

    @Override
    public CommonResult registerUser(RegisterInfo registerInfo,
                                     CommonRequestBody commonRequestBody) {
        // === Transaction codes ===
        List<String> children = new LinkedList<>();
        if (commonRequestBody.getChild() != null && !commonRequestBody.getChild().equals("")) {
            children.add(commonRequestBody.getChild());
        }
        children.add("userInfoService.createUserInfo");
        // === Transaction codes ===

        if (!registerMapper.insertUser(registerInfo, commonRequestBody, children)) {
            return new CommonResult(500, "insert user failed");
        }

        UserInfo userInfo = new UserInfo();
        userInfo.setUsername(registerInfo.getUsername());
        userInfo.setEmail(registerInfo.getEmail());
        userInfo.setPhoneNumber(registerInfo.getPhoneNumber());

        // === Transaction codes ===
        commonRequestBody.setParentUUID(commonRequestBody.getServiceUUID());
        commonRequestBody.setServiceUUID("userInfoService.createUserInfo");
        commonRequestBody.setChild("profileService.createProfile");
        // === Transaction codes ===

        CommonResult res = userInfoService.createUserInfo(userInfo, commonRequestBody.getGlobalTransactionUUID(),
                commonRequestBody.getServiceUUID(), commonRequestBody.getParentUUID(), commonRequestBody.getChild());
        if (!(res.getCode() > 199 && res.getCode() < 300)) {
            return res;
        }

        Profile profile = new Profile();
        profile.setUsername(registerInfo.getUsername());

        // === Transaction codes ===
        commonRequestBody.setParentUUID(commonRequestBody.getServiceUUID());
        commonRequestBody.setServiceUUID("profileService.createProfile");
        commonRequestBody.setChild("");
        // === Transaction codes ===

        res = profileService.createProfile(profile, commonRequestBody.getGlobalTransactionUUID(),
                commonRequestBody.getServiceUUID(), commonRequestBody.getParentUUID(), commonRequestBody.getChild());
        if (!(res.getCode() > 199 && res.getCode() < 300)) {
            return res;
        }

        res.setCode(200);
        res.setMessage("register success");
        return res;
    }
}
