package edu.bridge.domain;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.UUID;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class CommonRequestBody<T> {
    private UUID globalTransactionUUID;
    private String serviceUUID;
    private String parentUUID;
}
