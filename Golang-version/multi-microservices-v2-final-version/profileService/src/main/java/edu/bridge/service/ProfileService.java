package edu.bridge.service;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.CommonResult;
import edu.bridge.domain.Profile;

public interface ProfileService {
    CommonResult createProfile(Profile profile,
                               CommonRequestBody commonRequestBody);
}
