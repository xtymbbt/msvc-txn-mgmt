package edu.bridge.service;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.CommonResult;
import edu.bridge.domain.RegisterInfo;

public interface RegisterService {
    CommonResult registerUser(RegisterInfo registerInfo,
                              CommonRequestBody commonRequestBody);
}
