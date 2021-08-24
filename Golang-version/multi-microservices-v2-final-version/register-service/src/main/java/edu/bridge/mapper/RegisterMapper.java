package edu.bridge.mapper;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.RegisterInfo;

import java.util.HashMap;
import java.util.List;

public interface RegisterMapper {
    boolean updateUser();
    boolean insertUser(RegisterInfo registerInfo,
                       CommonRequestBody commonRequestBody,
                       List<String> children);
    boolean deleteUser();
    RegisterInfo queryUser();
}
