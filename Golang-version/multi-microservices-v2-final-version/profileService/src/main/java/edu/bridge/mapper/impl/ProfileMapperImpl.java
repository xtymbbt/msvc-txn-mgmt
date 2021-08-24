package edu.bridge.mapper.impl;

import edu.bridge.client.ExecTxnRpcGrpcClient;
import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.Profile;
import edu.bridge.mapper.ProfileMapper;
import edu.bridge.tools.CommonTools;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;
import java.lang.reflect.Field;
import java.util.HashMap;
import java.util.List;

@Service
public class ProfileMapperImpl implements ProfileMapper {
    @Value("${gRPC.dbName}")
    private String dbName;

    @Autowired
    private ExecTxnRpcGrpcClient grpcClient;

    @Override
    public boolean insertProfile(Profile profile, CommonRequestBody commonRequestBody, List<String> children) {
        // =====↓ ↓ ↓ ↓ ↓===== We can write this into a Spring Annotation.=====↓ ↓ ↓ ↓ ↓=====
        HashMap<String, String> data = new HashMap<>();
        // loop registerInfo's all fields through Java's reflection.
        Class<? extends Profile> cls = profile.getClass();
        Field[] fields = cls.getDeclaredFields();
        for (Field f : fields) {
            f.setAccessible(true);
            try {
                Object value = f.get(profile);
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
                "profile", true, true, "", data);
        // =====↑ ↑ ↑ ↑ ↑===== We can write this into a Spring Annotation.=====↑ ↑ ↑ ↑ ↑=====
    }
}
