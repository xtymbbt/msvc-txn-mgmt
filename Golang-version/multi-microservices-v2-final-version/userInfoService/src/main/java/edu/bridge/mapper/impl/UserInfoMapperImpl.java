package edu.bridge.mapper.impl;

import edu.bridge.client.CommonInfoGrpcClient;
import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.UserInfo;
import edu.bridge.mapper.UserInfoMapper;
import edu.bridge.tools.CommonTools;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import java.lang.reflect.Field;
import java.util.HashMap;

@Service
public class UserInfoMapperImpl implements UserInfoMapper {
    @Value("${gRPC.dbName}")
    private String dbName;

    @Autowired
    private CommonInfoGrpcClient grpcClient;

    @Override
    public boolean insertUserInfo(UserInfo userInfo, CommonRequestBody commonRequestBody, HashMap<String, Boolean> children) {
        // =====↓ ↓ ↓ ↓ ↓===== We can write this into a Spring Annotation.=====↓ ↓ ↓ ↓ ↓=====
        HashMap<String, String> data = new HashMap<>();
        // loop registerInfo's all fields through Java's reflection.
        Class<? extends UserInfo> cls = userInfo.getClass();
        Field[] fields = cls.getDeclaredFields();
        for (Field f : fields) {
            f.setAccessible(true);
            try {
                Object value = f.get(userInfo);
                if (value != null) {
                    System.out.println(value.getClass());
                    if (value.getClass() == String.class) {
                        value = "\"" + value + "\"";
                        System.out.println(value);
                    }
                    data.put(CommonTools.humpToLine(f.getName()), value.toString());
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
        return grpcClient.sendToDataCenter(true, commonRequestBody, children, dbName,
                "user_info", true, true, "", data);
        // =====↑ ↑ ↑ ↑ ↑===== We can write this into a Spring Annotation.=====↑ ↑ ↑ ↑ ↑=====
    }
}
