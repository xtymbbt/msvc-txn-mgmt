package com.bridge.service;

import com.bridge.grpc.SimpleClient;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/14 19:49
 */
@Service
public class HelloServiceImpl implements IHelloService{
    private Logger logger = LoggerFactory.getLogger(HelloServiceImpl.class);
    @Value("${gRPC.host}")
    private String host;
    @Value("${gRPC.port}")
    private int port;

    @Override
    public String sayHello(String name) {
        SimpleClient client = new SimpleClient(host,port);
        String replay = client.sayHello(name);
        try {
            client.shutdown();
        } catch (InterruptedException e) {
            logger.error("channel关闭异常：err={}",e.getMessage());
        }
        return replay;
    }

}