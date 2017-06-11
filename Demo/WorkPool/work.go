/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/6/11 5:58
 */

package WorkPool

type WorkRequest struct {
	Execute func(config interface{}) error
}
