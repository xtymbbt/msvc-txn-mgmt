package edu.bridge.domain;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class UserInfo {
    private String username;
    private Integer phoneNumber;
    private String email;
    private String photo;
    private Float longitude;
    private Float latitude;
    private String currentAddress;
    private String hobbyTags;
    private String consumeTags;
}
