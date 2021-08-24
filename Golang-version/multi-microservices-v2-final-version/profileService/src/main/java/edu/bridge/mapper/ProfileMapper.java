package edu.bridge.mapper;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.Profile;

import java.util.HashMap;
import java.util.List;

public interface ProfileMapper {
    boolean insertProfile(Profile profile,
                          CommonRequestBody commonRequestBody,
                          List<String> children);
}
