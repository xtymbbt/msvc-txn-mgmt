package edu.bridge.domain;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.UUID;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class CommonRequestBody {
    private UUID globalTransactionUUID;
    private String serviceUUID;
    private String parentUUID;
    private String child;
}
