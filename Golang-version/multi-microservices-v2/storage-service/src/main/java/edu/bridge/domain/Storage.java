package edu.bridge.domain;

import lombok.Data;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/16 13:41
 */
@Data
public class Storage {
    private Long id;
    private Long productId;
    private Integer total;
    private Integer used;
    private Integer residue;
}
