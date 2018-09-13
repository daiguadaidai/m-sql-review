// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/daiguadaidai/m-sql-review/config"
	"github.com/daiguadaidai/m-sql-review/service"
	"github.com/outbrain/golib/log"
)

var runConfig *config.ReviewConfig

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "m-sql-review",
	Short: "SQL审核工具",
	Long: `
    一款SQL审核工具, 主要用于MySQL SQL 相关审核. 启动工具后会提供一个http接口为用户实时链接并且审核相关SQL.
    启动工具:
    ./m-sql-review \
        --rule-name-length=100 \
        --rule-name-reg="^[a-zA-Z\$_][a-zA-Z\$\d_]*$" \
        --rule-charset="utf8,utf8mb4" \
        --rule-collate="utf8_general_ci,utf8mb4_general_ci" \
        --rule-allow-drop-database=false \
        --rule-allow-drop-table=false \
        --rule-allow-rename-table=false \
        --rule-allow-truncate-table=false
        --rule-table-engine="innodb"
    `,
	Run: func(cmd *cobra.Command, args []string) {
		err := service.Run(runConfig)
		if err != nil {
			log.Errorf("运行失败: %v", err)
		}
		fmt.Println(runConfig)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	runConfig = new(config.ReviewConfig)

	rootCmd.Flags().IntVar(&runConfig.RuleNameLength, "rule-name-length",
		config.RULE_NAME_LENGTH,"通用名称长度")
	rootCmd.Flags().StringVar(&runConfig.RuleNameReg, "rule-name-reg",
		config.RULE_NAME_REG, "通用名称匹配规则")
	rootCmd.Flags().StringVar(&runConfig.RuleCharSet, "rule-charset",
		config.RULE_CHARSET,"通用允许的字符集, 默认(多个用逗号隔开)")
	rootCmd.Flags().StringVar(&runConfig.RuleCollate, "rule-collate",
		config.RULE_COLLATE,"通用允许的collate, 默认(多个用逗号隔开)")
	rootCmd.Flags().BoolVar(&runConfig.RuleAllowDropDatabase, "rule-allow-drop-database",
		config.RULE_ALLOW_DROP_DATABASE,
		fmt.Sprintf("是否允许删除数据库, 默认: %v", config.RULE_ALLOW_DROP_DATABASE))
	rootCmd.Flags().BoolVar(&runConfig.RuleAllowDropTable, "rule-allow-drop-table",
		config.RULE_ALLOW_DROP_TABLE,
		fmt.Sprintf("是否允许删除表, 默认: %v", config.RULE_ALLOW_DROP_TABLE))
	rootCmd.Flags().BoolVar(&runConfig.RuleAllowRenameTable, "rule-allow-rename-table",
		config.RULE_ALLOW_RENAME_TABLE,
		fmt.Sprintf("是否允许重命名表, 默认: %v", config.RULE_ALLOW_RENAME_TABLE))
	rootCmd.Flags().BoolVar(&runConfig.RuleAllowTruncateTable, "rule-allow-truncate-table",
		config.RULE_ALLOW_TRUNCATE_TABLE,
		fmt.Sprintf("是否允许truncate表, 默认: %v", config.RULE_ALLOW_TRUNCATE_TABLE))
	rootCmd.Flags().StringVar(&runConfig.RuleTableEngine, "rule-table-engine",
		config.RULE_TABLE_ENGINE, "允许的存储引擎 默认(多个用逗号隔开)")
	rootCmd.Flags().StringVar(&runConfig.RuleTableEngine, "rule-not-allow-column-type",
		config.RULE_NOT_ALLOW_COLUMN_TYPE,
		"不允许的字段类型, 至此的类型: " +
		"decimal, tinyint, smallint, int, float, double, timestamp, bigint, mediumint, date, time, " +
		"datetime, year, newdate, varchar, bit, json, newdecimal, enum, set, tinyblob, mediumblob, " +
		"longblob, blob, tinytext, mediumtext, longtext, text, geometry")
	rootCmd.Flags().BoolVar(&runConfig.RuleNeedTableComment, "rule-need-table-comment",
		config.RULE_NEED_TABLE_COMMENT,
		fmt.Sprintf("表是否需要注释 默认: %v", config.RULE_NEED_TABLE_COMMENT))
	rootCmd.Flags().BoolVar(&runConfig.RuleNeedColumnComment, "rule-need-column-comment",
		config.RULE_NEED_COLUMN_COMMENT,
		fmt.Sprintf("字段是否需要注释 默认: %v", config.RULE_NEED_COLUMN_COMMENT))
	rootCmd.Flags().BoolVar(&runConfig.RulePKAutoIncrement, "rule-pk-auto-increment",
		config.RULE_PK_AUTO_INCREMENT,
		fmt.Sprintf("主键字段中是否需要有自增字段 默认: %v", config.RULE_PK_AUTO_INCREMENT))
	rootCmd.Flags().BoolVar(&runConfig.RuleNeedPK, "rule-need-pk",
		config.RULE_NEED_PK,
		fmt.Sprintf("建表是否需要主键 默认: %v", config.RULE_NEED_PK))
	rootCmd.Flags().IntVar(&runConfig.RuleIndexColumnCount, "rule-index-column-count",
		config.RULE_INDEX_COLUMN_COUNT, "索引允许字段个数")
	rootCmd.Flags().StringVar(&runConfig.RuleTableNameReg, "rule-table-name-reg",
		config.RULE_TABLE_NAME_GRE, "表名, 名命名规范(正则)")
	rootCmd.Flags().StringVar(&runConfig.RuleIndexNameReg, "rule-index-name-reg",
		config.RULE_INDEX_NAME_REG, "索引名命名规范(正则)")
	rootCmd.Flags().StringVar(&runConfig.RuleUniqueIndexNameReg, "rule-unique-index-name-reg",
		config.RULE_UNIQUE_INDEX_NAME_REG, "唯一索引名命名规范(正则)")
	rootCmd.Flags().BoolVar(&runConfig.RuleAllColumnNotNull, "rule-all-column-not-null",
		config.RULE_ALL_COLUMN_NOT_NULL,
		fmt.Sprintf("是否所有字段. 默认: %v", config.RULE_ALL_COLUMN_NOT_NULL))
	rootCmd.Flags().BoolVar(&runConfig.RuleAllowForeignKey, "rule-allow-foreign-key",
		config.RULE_ALLOW_FOREIGN_KEY,
		fmt.Sprintf("是否允许使用外键. 默认: %v", config.RULE_ALLOW_FOREIGN_KEY))
	rootCmd.Flags().BoolVar(&runConfig.RuleAllowFullText, "rule-allow-full-text",
		config.RULE_ALLOW_FULL_TEXT,
		fmt.Sprintf("是否允许使用全文索引. 默认: %v", config.RULE_ALLOW_FULL_TEXT))
	rootCmd.Flags().StringVar(&runConfig.RuleNotNullColumnType, "rule-not-null-column-type",
		config.RULE_NOT_NULL_COLUMN_TYPE,
		"必须为not null的字段类型, 默认(多个用逗号隔开). 可填写的类型有: " +
		"decimal, tinyint, smallint, int, float, double, timestamp, bigint, mediumint, date, time, " +
		"datetime, year, newdate, varchar, bit, json, newdecimal, enum, set, tinyblob, mediumblob, " +
		"longblob, blob, tinytext, mediumtext, longtext, text, geometry")
	rootCmd.Flags().IntVar(&runConfig.RuleTextTypeColumnCount, "rule-text-type-column-count",
		config.RULE_TEXT_TYPE_COLUMN_COUNT, "允许使用text/blob字段个数. 如果在rule-not-allow-column-type相关text字段." +
		"该参数将不其作用")
	rootCmd.Flags().StringVar(&runConfig.RuleNeedIndexColumnName, "rule-need-index-column-name",
		config.RULE_NEED_INDEX_COLUMN_NAME, "必须要有索引的字段名, 默认(多个用逗号隔开)")
}
