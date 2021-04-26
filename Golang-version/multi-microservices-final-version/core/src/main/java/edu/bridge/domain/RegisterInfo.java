package edu.bridge.domain;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class RegisterInfo {
    private String username;
    private String password;
    private Integer phoneNumber;
    private String email;
}
