package edu.bridge.mapper.impl;

import edu.bridge.client.CommonInfoGrpcClient;
import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.RegisterInfo;
import edu.bridge.mapper.RegisterMapper;
import edu.bridge.tools.CommonTools;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import java.lang.reflect.Field;
import java.util.HashMap;

@Service
public class RegisterMapperImpl implements RegisterMapper {
    @Value("${gRPC.dbName}")
    private String dbName;

    @Autowired
    private CommonInfoGrpcClient grpcClient;

    @Override
    public boolean updateUser() {
        return false;
    }

    @Override
    public boolean insertUser(RegisterInfo registerInfo,
                              CommonRequestBody commonRequestBody,
                              HashMap<String, Boolean> children) {
        // =====↓ ↓ ↓ ↓ ↓===== We can write this into a Spring Annotation.=====↓ ↓ ↓ ↓ ↓=====
        HashMap<String, String> data = new HashMap<>();
        // loop registerInfo's all fields through Java's reflection.
        Class<? extends RegisterInfo> cls = registerInfo.getClass();
        Field[] fields = cls.getDeclaredFields();
        for (Field f : fields) {
            f.setAccessible(true);
            try {
                Object value = f.get(registerInfo);
                if (value != null) {
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
                "register_info", true, true, "", data);
        // =====↑ ↑ ↑ ↑ ↑===== We can write this into a Spring Annotation.=====↑ ↑ ↑ ↑ ↑=====
    }

    @Override
    public boolean deleteUser() {
        return false;
    }

    @Override
    public RegisterInfo queryUser() {
        return null;
    }
}
