package edu.bridge.service.impl;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.CommonResult;
import edu.bridge.domain.Profile;
import edu.bridge.mapper.ProfileMapper;
import edu.bridge.service.ProfileService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashMap;

@Service
public class ProfileServiceImpl implements ProfileService {
    @Autowired
    private ProfileMapper profileMapper;

    @Override
    public CommonResult createProfile(Profile profile,
                                      CommonRequestBody commonRequestBody) {
        // === Transaction codes ===
        HashMap<String, Boolean> children = new HashMap<>();
        if (commonRequestBody.getChild() != null && !commonRequestBody.getChild().equals("")) {
            children.put(commonRequestBody.getChild(), true);
        }
        // === Transaction codes ===

        return profileMapper.insertProfile(profile, commonRequestBody, children)
                ? new CommonResult(200, "insert userInfo success")
                : new CommonResult(500, "insert userInfo failed");
    }
}
