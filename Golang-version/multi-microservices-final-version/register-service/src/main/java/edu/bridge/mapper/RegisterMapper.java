package edu.bridge.mapper;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.RegisterInfo;

import java.util.HashMap;

public interface RegisterMapper {
    boolean updateUser();
    boolean insertUser(RegisterInfo registerInfo,
                       CommonRequestBody commonRequestBody,
                       HashMap<String, Boolean> children);
    boolean deleteUser();
    RegisterInfo queryUser();
}
