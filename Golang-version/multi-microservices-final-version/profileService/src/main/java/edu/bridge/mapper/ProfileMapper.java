package edu.bridge.mapper;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.Profile;

import java.util.HashMap;

public interface ProfileMapper {
    boolean insertProfile(Profile profile,
                          CommonRequestBody commonRequestBody,
                          HashMap<String, Boolean> children);
}
